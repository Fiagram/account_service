package database

import (
	"context"
	"fmt"

	"github.com/Fiagram/account_service/internal/utils"
	"go.uber.org/zap"
)

type AccountRole struct {
	Id   uint8  `json:"id"`
	Name string `json:"name"`
}

type AccountRoleAccessor interface {
	GetRoleById(ctx context.Context, id uint8) (AccountRole, error)
	GetRoleByName(ctx context.Context, name string) (AccountRole, error)
	WithExecutor(exec Executor) AccountRoleAccessor
}

type accountRoleAccessor struct {
	exec   Executor
	logger *zap.Logger
}

func NewAccountRoleAccessor(
	exec Executor,
	logger *zap.Logger,
) AccountRoleAccessor {
	return &accountRoleAccessor{
		exec:   exec,
		logger: logger,
	}
}

func (a accountRoleAccessor) GetRoleById(
	ctx context.Context,
	id uint8,
) (AccountRole, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("role_id", id))

	query := fmt.Sprintf(`SELECT id, name FROM account_role WHERE id = %d`, id)
	row := a.exec.QueryRowContext(ctx, query)

	var out AccountRole
	err := row.Scan(&out.Id, &out.Name)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account row by id")
		return AccountRole{}, err
	}

	return out, nil
}

func (a accountRoleAccessor) GetRoleByName(
	ctx context.Context,
	name string,
) (AccountRole, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("role_name", name))

	query := fmt.Sprintf(`SELECT id, name FROM account_role WHERE name = "%s"`, name)
	row := a.exec.QueryRowContext(ctx, query)

	var out AccountRole
	err := row.Scan(&out.Id, &out.Name)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get account row by name")
		return AccountRole{}, err
	}

	return out, nil
}

func (a accountRoleAccessor) WithExecutor(
	exec Executor,
) AccountRoleAccessor {
	return &accountRoleAccessor{
		exec:   exec,
		logger: a.logger,
	}
}
