package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Fiagram/account_service/internal/configs"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

type Executor interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)

	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)

	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row

	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func InitAndMigrateUpDatabase(databaseConfig configs.Database, logger *zap.Logger) (*sql.DB, func(), error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		databaseConfig.Username,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed when connecting the database")
		return nil, nil, err
	}

	cleanupDb := func() {
		db.Close()
	}

	migrator := NewMigrator(db, logger)
	err = migrator.Up(context.Background())
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to execute database migration up")
		cleanupDb()
		return nil, nil, err
	}

	return db, cleanupDb, nil
}
