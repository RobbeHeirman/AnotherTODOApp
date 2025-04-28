package postgres

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/robbeheirman/todo/auth/models"
)

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
	return nil
}

//go:embed sql/user_scheme.sql
var schemaSQL string

func (repo *Repository) Install() error {
	connStr := repo.getConnectionString()
	conn, err := pgx.Connect(context.Background(), connStr)
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
