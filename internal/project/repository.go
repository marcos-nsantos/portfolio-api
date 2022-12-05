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

func (r *Repo) FindAll(ctx context.Context) ([]*entity.Project, error) {
	var projects []*entity.Project
	err := r.db.WithContext(ctx).
		Model(&entity.Project{}).
		Select("id", "name", "description", "url").
		Find(&projects).
		Order("id desc").Error
	return projects, err
}

func (r *Repo) FindByID(ctx context.Context, id uint64) (*entity.Project, error) {
	var project entity.Project
	err := r.db.WithContext(ctx).
		Model(&entity.Project{}).
		Select("id", "name", "description", "url").
		First(&project, id).Error
	return &project, err
}

func (r *Repo) Update(ctx context.Context, project *entity.Project) error {
	return r.db.WithContext(ctx).
		Model(project).
		Select("name", "description", "url").
		Updates(project).Error
}

func (r *Repo) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&entity.Project{}, id).Error
}
