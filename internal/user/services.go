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

func (s *Services) GetAll(ctx context.Context) ([]*entity.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *Services) Update(ctx context.Context, user *entity.User) error {
	if _, err := s.GetByID(ctx, user.ID); err != nil {
		return err
	}
	return s.repo.Update(ctx, user)
}
