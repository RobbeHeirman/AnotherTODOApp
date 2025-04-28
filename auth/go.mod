module github.com/robbeheirman/todo/auth

go 1.24.2

require (
	github.com/robbeheirman/todo/shared v0.0.0
	golang.org/x/crypto v0.37.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.4 // indirect
	golang.org/x/text v0.24.0 // indirect
)

replace github.com/robbeheirman/todo/shared => ../shared
