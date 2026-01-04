package logic

import (
	"context"
	"database/sql"

	"github.com/Fiagram/account_service/internal/dataaccess/database"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Account interface {
	CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error)
	DeleteAccount(ctx context.Context, params DeleteAccountParams) error
	CheckAccountValid(ctx context.Context, params CheckAccountValidParams) (CheckAccountValidOutput, error)
	// CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionOutput, error)
}

type account struct {
	db                      *sql.DB
	accountAccessor         database.AccountAccessor
	accountPasswordAccessor database.AccountPasswordAccessor
	hashLogic               Hash
	logger                  *zap.Logger
}

func NewAccount(
	db *sql.DB,
	accountAccessor database.AccountAccessor,
	accountPasswordAccessor database.AccountPasswordAccessor,
	hashLogic Hash,
	logger *zap.Logger,
) Account {
	return &account{
		db:                      db,
		accountAccessor:         accountAccessor,
		accountPasswordAccessor: accountPasswordAccessor,
		hashLogic:               hashLogic,
		logger:                  logger,
	}
}

func (a account) CreateAccount(
	ctx context.Context,
	params CreateAccountParams,
) (CreateAccountOutput, error) {
	emptyOutput := CreateAccountOutput{}
	isUsernameTaken, err := a.accountAccessor.IsUsernameTaken(ctx, params.AccountInfo.Username)
	if err != nil {
		return emptyOutput, status.Error(codes.Internal, "failed to check if username taken")
	} else if isUsernameTaken {
		return emptyOutput, status.Error(codes.AlreadyExists, "username has already taken")
	}

	tx, err := a.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return emptyOutput, ErrTxInitFailed
	}
	defer tx.Rollback()

	id, err := a.accountAccessor.
		WithExecutor(tx).
		CreateAccount(ctx, database.Account{
			Username:    params.AccountInfo.Username,
			Fullname:    params.AccountInfo.Fullname,
			Email:       params.AccountInfo.Email,
			PhoneNumber: params.AccountInfo.PhoneNumber,
			RoleId:      uint8(params.AccountInfo.Role),
		})
	if err != nil {
		return emptyOutput, status.Error(codes.Internal, "failed to create new account")
	}

	hashedString, err := a.hashLogic.Hash(ctx, params.Password)
	if err != nil {
		return emptyOutput, err
	}

	err = a.accountPasswordAccessor.
		WithExecutor(tx).
		CreateAccountPassword(ctx, database.AccountPassword{
			OfAccountId:  id,
			HashedString: hashedString,
		})
	if err != nil {
		return emptyOutput, status.Error(codes.Internal, "failed to create new password")
	}

	if err = tx.Commit(); err != nil {
		return emptyOutput, ErrTxCommitFailed
	}

	return CreateAccountOutput{
		AccountId: id,
	}, nil
}

func (a account) DeleteAccount(
	ctx context.Context,
	params DeleteAccountParams,
) error {
	isExisted, err := a.accountAccessor.IsUsernameTaken(ctx, params.Username)
	if err != nil {
		return status.Error(codes.Internal, "failed to check if username existed")
	} else if !isExisted {
		return status.Error(codes.NotFound, "username does not existed")
	}

	tx, err := a.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return ErrTxInitFailed
	}
	defer tx.Rollback()
	acc, err := a.accountAccessor.
		WithExecutor(tx).
		GetAccountByUsername(ctx, params.Username)
	if err != nil {
		return status.Error(codes.Internal, "failed to get account")
	}
	err = a.accountPasswordAccessor.
		WithExecutor(tx).
		DeleteAccountPassword(ctx, acc.Id)
	if err != nil {
		return status.Error(codes.Internal, "failed to delete password")
	}
	err = a.accountAccessor.
		WithExecutor(tx).
		DeleteAccountById(ctx, acc.Id)
	if err != nil {
		return status.Error(codes.Internal, "failed to delete account")
	}

	if err = tx.Commit(); err != nil {
		return ErrTxCommitFailed
	}

	return nil
}

func (a account) CheckAccountValid(
	ctx context.Context,
	params CheckAccountValidParams,
) (CheckAccountValidOutput, error) {
	emptyObj := CheckAccountValidOutput{}
	acc, err := a.accountAccessor.
		GetAccountByUsername(ctx, params.Username)
	if err != nil {
		return emptyObj, status.Error(codes.NotFound, "failed to get account")
	}

	truly, err := a.accountPasswordAccessor.
		GetAccountPassword(ctx, acc.Id)
	if err != nil {
		return emptyObj, status.Error(codes.NotFound, "failed to get password")
	}

	isValid, err := a.hashLogic.
		IsHashEqual(ctx, params.Password, truly.HashedString)
	if err != nil {
		return emptyObj, status.Error(codes.Internal, "failed to when equal hashed")
	}

	if !isValid {
		return emptyObj, nil
	}
	return CheckAccountValidOutput{
		AccountId: acc.Id,
	}, nil
}
