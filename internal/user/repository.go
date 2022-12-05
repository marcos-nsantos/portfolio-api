package user

import (
	"context"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Insert(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *Repo) FindAll(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Select("id", "first_name", "last_name", "email").
		Find(&users).
		Order("id desc").Error
	return users, err
}

func (r *Repo) FindByID(ctx context.Context, id uint64) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).
		Model(&entity.User{}).
		Select("id", "first_name", "last_name", "email").
		First(&user, id).Error
	return &user, err
}

func (r *Repo) Update(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).
		Select("first_name", "last_name", "email").
		Updates(user).Error
}
