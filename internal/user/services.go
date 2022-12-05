package user

import (
	"context"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
)

type Services struct {
	repo Repository
}

func NewServices(repo Repository) *Services {
	return &Services{repo: repo}
}

func (s *Services) Create(ctx context.Context, user *entity.User) error {
	return s.repo.Insert(ctx, user)
}

func (s *Services) GetByID(ctx context.Context, id uint64) (*entity.User, error) {
	return s.repo.FindByID(ctx, id)
}
