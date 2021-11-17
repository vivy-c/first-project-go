package services

import (
	"github.com/vivy-c/first-project-go/internal/repository"
)

type service struct {
	mysqlrepo repository.Repository
}

func NewService(mysqlrepo repository.Repository) Services {
	return &service{mysqlrepo}
}
