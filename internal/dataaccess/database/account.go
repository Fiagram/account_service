package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Fiagram/account_service/internal/utils"
	"go.uber.org/zap"
)

type Account struct {
	Id          uint64    `sql:"id"`
	Username    string    `sql:"username"`
	Fullname    string    `sql:"fullname"`
	Email       string    `sql:"email"`
	PhoneNumber string    `sql:"phone_number"`
	RoleId      uint8     `sql:"role_id"`
	CreatedAt   time.Time `sql:"created_at"`
	UpdatedAt   time.Time `sql:"updated_at"`
}

type AccountAccessor interface {
	CreateAccount(ctx context.Context, account Account) (uint64, error)
	// GetAccountById(ctx context.Context, id uint64) (Account, error)
	// GetAccountByUsername(ctx context.Context, username string) (Account, error)
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

func (a accountAccessor) CreateAccount(ctx context.Context, acc Account) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account", acc))

	query := fmt.Sprintf(`INSERT INTO accounts 
			(username, fullname, email, phone_number, role_id)
			VALUES (%s, %s, %s, %s, %d)`,
		acc.Username,
		acc.Fullname,
		acc.Email,
		acc.PhoneNumber,
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

func (a accountAccessor) WithExecutor(exec Executor) AccountAccessor {
	return &accountAccessor{
		exec: exec,
	}
}

// func (a accountAccessor) GetAccountById(ctx context.Context, id uint64) (Account, error) {}
// func (a accountAccessor) GetAccountByUsername(ctx context.Context, username string) (Account, error) {
// }
