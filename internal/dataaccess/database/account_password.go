package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Fiagram/account_service/internal/utils"
	"go.uber.org/zap"
)

type AccountPassword struct {
	OfAccountId  uint64    `json:"of_account_id"`
	HashedString string    `json:"hashed_string"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AccountPasswordAccessor interface {
	CreateAccountPassword(ctx context.Context, ap AccountPassword) error
	GetAccountPassword(ctx context.Context, id uint64) (AccountPassword, error)
	UpdateAccountPassword(ctx context.Context, ap AccountPassword) error
	DeleteAccountPassword(ctx context.Context, id uint64) error
	WithExecutor(exec Executor) AccountPasswordAccessor
}

type accountPasswordAccessor struct {
	exec   Executor
	logger *zap.Logger
}

func NewAccountPasswordAccessor(
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
) error {
	if ap.OfAccountId == 0 {
		return fmt.Errorf("lack of information")
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("of_account_id", ap.OfAccountId))
	if ap.OfAccountId == 0 {
		return fmt.Errorf("lack of information")
	}
	query := fmt.Sprintf(`INSERT INTO account_passwords 
		(of_account_id, hashed_string)
		VALUES ("%d", "%s")`,
		ap.OfAccountId,
		strings.TrimSpace(ap.HashedString),
	)

	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to insert account password")
		return err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return err
	}

	return nil
}

func (a accountPasswordAccessor) GetAccountPassword(
	ctx context.Context,
	id uint64,
) (AccountPassword, error) {
	if id == 0 {
		return AccountPassword{}, fmt.Errorf("lack of information")
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("of_account_id", id))
	query := fmt.Sprintf(`SELECT * 
			FROM account_passwords 
			WHERE of_account_id = "%d"`, id)
	row := a.exec.QueryRowContext(ctx, query)
	var out AccountPassword
	err := row.Scan(&out.OfAccountId,
		&out.HashedString,
		&out.CreatedAt,
		&out.UpdatedAt)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get password")
		return AccountPassword{}, err
	}

	return out, nil
}

func (a accountPasswordAccessor) DeleteAccountPassword(
	ctx context.Context,
	id uint64,
) error {
	if id == 0 {
		return fmt.Errorf("lack of information")
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("of_account_id", id))
	if id == 0 {
		return fmt.Errorf("lack of information")
	}
	query := fmt.Sprintf(`DELETE FROM account_passwords 
			WHERE of_account_id = "%d"`, id)
	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to delete password")
		return err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return err
	}

	return nil
}

func (a accountPasswordAccessor) UpdateAccountPassword(
	ctx context.Context,
	ap AccountPassword,
) error {
	if ap.OfAccountId == 0 && ap.HashedString == "" {
		return fmt.Errorf("lack of information")
	}

	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Any("of_account_id", ap.OfAccountId))
	query := fmt.Sprintf(`UPDATE account_passwords SET 
			hashed_string = "%s" 
			WHERE of_account_id = "%d"`,
		strings.TrimSpace(ap.HashedString),
		ap.OfAccountId,
	)

	result, err := a.exec.ExecContext(ctx, query)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to update password")
		return err
	}

	rowEfNum, err := result.RowsAffected()
	if rowEfNum != 1 && err != nil {
		logger.With(zap.Error(err)).Error("failed to effect row")
		return err
	}

	return nil
}

func (a accountPasswordAccessor) WithExecutor(
	exec Executor,
) AccountPasswordAccessor {
	return &accountPasswordAccessor{
		exec:   exec,
		logger: a.logger,
	}
}
