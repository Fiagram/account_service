package database

import (
	"context"
	"fmt"

	"github.com/Fiagram/account_service/internal/utils"
	"go.uber.org/zap"
)

type AccountPassword struct {
	OfAccountId  uint64
	HashedString string
}

type AccountPasswordAccessor interface {
	CreateAccountPassword(ctx context.Context, ap AccountPassword) (uint64, error)
	WithExecutor(exec Executor) AccountPasswordAccessor
}

type accountPasswordAccessor struct {
	exec   Executor
	logger *zap.Logger
}

func NewAccountPassword(
	exec Executor,
	logger *zap.Logger,
) AccountPasswordAccessor {
	return &accountPasswordAccessor{
		exec:   exec,
		logger: logger,
	}
}

func (a accountPasswordAccessor) CreateAccountPassword(
	ctx context.Context,
	ap AccountPassword,
) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("account_password", ap))

	query := fmt.Sprintf(`INSERT INTO accounts (username, fullname, email, phone_number, role_id)
							VALUES (%s, %s, %s, %s, %d)`)

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

func (a accountPasswordAccessor) WithExecutor(
	exec Executor,
) AccountPasswordAccessor {
	return &accountPasswordAccessor{
		exec:   exec,
		logger: a.logger,
	}
}
