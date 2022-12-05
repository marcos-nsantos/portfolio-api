package user

import (
	"context"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type Writer interface {
	Insert(ctx context.Context, user *entity.User) error
}

type Reader interface {
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByID(ctx context.Context, id uint64) (*entity.User, error)
}

type Repository interface {
	Writer
	Reader
}

type Service interface{}
