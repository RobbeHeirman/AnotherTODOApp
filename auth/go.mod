module github.com/robbeheirman/todo/auth

go 1.24.2

require (
	github.com/robbeheirman/todo/shared v0.0.0
	golang.org/x/crypto v0.37.0
)

replace github.com/robbeheirman/todo/shared => ../shared
