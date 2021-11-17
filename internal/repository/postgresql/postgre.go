package repository

import (
	"github.com/vivy-c/first-project-go/internal/repository"

	"github.com/jmoiron/sqlx"
)

const ()

//buat ngehindarin sql injection
var statement PreparedStatement

type PreparedStatement struct {
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	// InitPreparedStatement(repo)
	return repo
}
