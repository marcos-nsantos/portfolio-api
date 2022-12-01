package project

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

func (r *Repo) Insert(ctx context.Context, project *entity.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}
