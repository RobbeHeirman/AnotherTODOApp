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

//go:embed sql/select_password.sql
var selectPassword string

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

func (repo *Repository) CreateUser(user *models.User) (*models.UserId, error) {
	connection, err := repo.getConnection()
	if err != nil {
		return nil, err
	}
	defer func(connection *pgx.Conn, ctx context.Context) {
		err := connection.Close(ctx)
		if err != nil {
			logger.Error(err.Error())
		}
	}(connection, context.Background())
	users, err := postgres.InsertAndGetObjects(connection, UserTable, rowToUser, user)
	if err != nil {
		return nil, err
	}
	if len(users) != 1 {
		return nil, errors.New("insert failed")
	}
	logger.Info("Inserted user")
	return users[0], nil
}

func (repo *Repository) GetUserByEmail(user *models.User) (models.UserLogsInDb, error) {
	connection, err := repo.getConnection()
	defer connection.Close(context.Background())
	if err != nil {
		return models.UserLogsInDb{}, err
	}

	rows, err := connection.Query(context.Background(), selectPassword, user.Email)
	defer rows.Close()
	if rows == nil || !rows.Next() {
		return models.UserLogsInDb{}, errors.New("User not found")
	}
	var id int
	var password string
	err = rows.Scan(&id, &password)
	if err != nil {
		return models.UserLogsInDb{}, err
	}
	return models.UserLogsInDb{
		Id:       id,
		Password: password,
	}, nil
}

func (repo *Repository) Install() error {
	conn, err := repo.getConnection()
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			logger.Error(err.Error())
		}
	}(conn, context.Background())
	if err != nil {
		logger.Error("Error db conn", err)
		return err
	}

	_, err = conn.Exec(context.Background(), schemaSQL)
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}
	logger.Info("Schema created")
	return nil
}

func (repo *Repository) getConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", repo.username, repo.password, repo.host, repo.port, repo.database)
}

func (repo *Repository) getConnection() (*pgx.Conn, error) {
	connStr := repo.getConnectionString()
	connect, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		logger.Error("Could not connect to db", err)
	}
	return connect, err
}

func rowToUser(row pgx.CollectableRow) (*models.UserId, error) {
	var id int
	err := row.Scan(&id)
	fmt.Printf("Called Row to User and id %v\n", id)
	if err != nil {
		logger.Error("Could not scan row", err)
		return nil, err
	}
	return &models.UserId{
		Id: id,
	}, nil
}
