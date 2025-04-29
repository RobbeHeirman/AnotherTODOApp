package postgres

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/robbeheirman/todo/auth/models"
	"github.com/robbeheirman/todo/shared/persistence/postgres"
	"log/slog"
)

var logger = slog.Default()

const UserTable = "users"

//go:embed sql/user_scheme.sql
var schemaSQL string

type Repository struct {
	host     string
	port     int
	database string
	username string
	password string
}

func NewRepository(host string, port int, database string, username string, password string) *Repository {
	return &Repository{
		host:     host,
		port:     port,
		database: database,
		username: username,
		password: password,
	}
}

func (repo *Repository) CreateUser(user *models.User) error {
	connection, err := repo.getConnection()
	if err != nil {
		return err
	}
	qty, err := postgres.InsertObject(connection, UserTable, user)
	if err != nil {
		return err
	}
	if qty != 1 {
		return errors.New("insert failed")
	}
	logger.Info("Inserted user")
	return nil
}

func (repo *Repository) Install() error {
	conn, err := repo.getConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), schemaSQL)
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}
	return nil
}

func (repo *Repository) getConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", repo.username, repo.password, repo.host, repo.port, repo.database)
}

func (repo *Repository) getConnection() (*pgx.Conn, error) {
	connStr := repo.getConnectionString()
	return pgx.Connect(context.Background(), connStr)

}
