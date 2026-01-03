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

	UpdateAccount(ctx context.Context, account Account) error

	DeleteAccountById(ctx context.Context, id uint64) error
	DeleteAccountByUsername(ctx context.Context, username string) error

	IsUsernameTaken(ctx context.Context, username string) (bool, error)

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
	if acc.Username == "" &&
		acc.RoleId == 0 {
		return 0, ErrLackOfInfor
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account", acc))
	query := fmt.Sprintf(`INSERT INTO accounts 
			(username, fullname, email, phone_number, role_id)
			VALUES ("%s", "%s", "%s", "%s", "%d")`,
		strings.TrimSpace(acc.Username),
		strings.TrimSpace(acc.Fullname),
		strings.TrimSpace(acc.Email),
		strings.TrimSpace(acc.PhoneNumber),
		acc.RoleId)

	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to create account")
		return 0, err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
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
	if id == 0 {
		return Account{}, ErrLackOfInfor
	}

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
	if username == "" {
		return Account{}, ErrLackOfInfor
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("username", username))
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
	if id == 0 {
		return ErrLackOfInfor
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account_id", id))
	query := fmt.Sprintf(`DELETE FROM accounts 
			WHERE id = "%d"`, id)
	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to delete account")
		return err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return err
	}

	return nil
}

func (a accountAccessor) DeleteAccountByUsername(
	ctx context.Context,
	username string,
) error {
	if username == "" {
		return ErrLackOfInfor
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("username", username))
	query := fmt.Sprintf(`DELETE FROM accounts 
			WHERE username = "%s"`, username)
	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to delete account")
		return err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return err
	}

	return nil
}

func (a accountAccessor) UpdateAccount(
	ctx context.Context,
	acc Account,
) error {
	if acc.Username == "" {
		return ErrLackOfInfor
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account", acc))
	query := fmt.Sprintf(`
			UPDATE accounts SET 
			fullname = "%s", 
			email = "%s", 
			phone_number = "%s", 
			role_id = "%d" 
			WHERE username = "%s"`,
		strings.TrimSpace(acc.Fullname),
		strings.TrimSpace(acc.Email),
		strings.TrimSpace(acc.PhoneNumber),
		acc.RoleId,
		strings.TrimSpace(acc.Username),
	)

	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to update account")
		return err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return err
	}

	return nil
}

func (a accountAccessor) IsUsernameTaken(
	ctx context.Context,
	username string,
) (bool, error) {
	if username == "" {
		return false, ErrLackOfInfor
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("username", username))
	const query = `SELECT EXISTS(SELECT 1 FROM accounts WHERE username = ?) AS is_taken`
	var isTaken int
	err := a.exec.QueryRowContext(ctx, query, username).Scan(&isTaken)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to check username taken")
		return false, err
	}

	if isTaken == 1 {
		return true, nil
	}
	return false, nil
}

func (a accountAccessor) WithExecutor(
	exec Executor,
) AccountAccessor {
	return &accountAccessor{
		exec:   exec,
		logger: a.logger,
	}
}
