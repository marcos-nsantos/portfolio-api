package user

import (
	"context"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type Writer interface {
	Insert(ctx context.Context, user *entity.User) error
}

type Reader interface{}

type Repository interface {
	Writer
	Reader
}

type Service interface{}
