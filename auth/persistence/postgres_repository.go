package persistence

type PostgresRepository struct{}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{}
}

func (repo *PostgresRepository) CreateUser(user *User) error {
	return nil
}
