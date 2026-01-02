package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Fiagram/account_service/internal/utils"
	"go.uber.org/zap"
)

type Account struct {
	Id          uint64    `json:"id"`
	Username    string    `json:"username"`
	Fullname    string    `json:"fullname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	RoleId      uint8     `json:"role_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AccountAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)

	GetAccountById(ctx context.Context, id uint64) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)

	// UpdateAccount(ctx context.Context, account Account) (uint64, error)

	DeleteAccountById(ctx context.Context, id uint64) error
	DeleteAccountByUsername(ctx context.Context, username string) error

	WithExecutor(exec Executor) AccountAccessor
}

type accountAccessor struct {
	exec   Executor
	logger *zap.Logger
}

func NewAccountAccessor(
	exec Executor,
	logger *zap.Logger,
) AccountAccessor {
	return &accountAccessor{
		exec:   exec,
		logger: logger,
	}
}

func (a accountAccessor) CreateAccount(
	ctx context.Context,
	acc Account,
) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account", acc))

	query := fmt.Sprintf(`INSERT INTO accounts 
			(username, fullname, email, phone_number, role_id)
			VALUES ("%s", "%s", "%s", "%s", "%d")`,
		strings.TrimSpace(acc.Username),
		strings.TrimSpace(acc.Fullname),
		strings.TrimSpace(acc.Email),
		strings.TrimSpace(acc.PhoneNumber),
		acc.RoleId)

	ressult, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create account")
		return 0, err
	}

	lastInsertedId, err := ressult.LastInsertId()
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get last inserted id")
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (a accountAccessor) GetAccountById(
	ctx context.Context,
	id uint64,
) (Account, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account_id", id))

	query := fmt.Sprintf(`SELECT * FROM accounts 
			WHERE id = "%d"`, id)
	row := a.exec.QueryRowContext(ctx, query)

	var out Account
	err := row.Scan(&out.Id,
		&out.Username,
		&out.Fullname,
		&out.Email,
		&out.PhoneNumber,
		&out.RoleId,
		&out.CreatedAt,
		&out.UpdatedAt)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account by id")
		return Account{}, err
	}

	return out, nil
}

func (a accountAccessor) GetAccountByUsername(
	ctx context.Context,
	username string,
) (Account, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account_username", username))

	query := fmt.Sprintf(`SELECT * FROM accounts 
			WHERE username = "%s"`, username)
	row := a.exec.QueryRowContext(ctx, query)

	var out Account
	err := row.Scan(&out.Id,
		&out.Username,
		&out.Fullname,
		&out.Email,
		&out.PhoneNumber,
		&out.RoleId,
		&out.CreatedAt,
		&out.UpdatedAt)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account by username")
		return Account{}, err
	}

	return out, nil
}

func (a accountAccessor) DeleteAccountById(
	ctx context.Context,
	id uint64,
) error {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account_id", id))
	query := fmt.Sprintf(`DELETE FROM accounts 
			WHERE id = "%d"`, id)
	_, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to delete account by id")
		return err
	}
	return nil
}

func (a accountAccessor) DeleteAccountByUsername(
	ctx context.Context,
	username string,
) error {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account_username", username))
	query := fmt.Sprintf(`DELETE FROM accounts 
			WHERE username = "%s"`, username)
	_, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to delete account by username")
		return err
	}
	return nil
}

func (a accountAccessor) WithExecutor(exec Executor) AccountAccessor {
	return &accountAccessor{
		exec: exec,
	}
}
