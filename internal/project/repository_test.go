//go:build integration

package project

import (
	"context"
	"fmt"
	"github.com/marcos-nsantos/portfolio-api/internal/user"
	"os"
	"testing"

	"github.com/marcos-nsantos/portfolio-api/internal/database"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	conn, err := database.New()
	if err != nil {
		fmt.Println(err)
	}
	db = conn.Client
	if err := conn.CreateTables(); err != nil {
		fmt.Println(err)
	}
	code := m.Run()
	if err := conn.DropTables(); err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func TestInsert(t *testing.T) {
	userEntity := &entity.User{
		FirstName: "Marcos",
		LastName:  "Santos",
		Email:     "email@email.com",
		Password:  "password",
	}
	useRepo := user.NewRepo(db)
	err := useRepo.Insert(context.Background(), userEntity)
	assert.NoError(t, err)

	project := &entity.Project{
		Name:        "Test",
		Description: "Test",
		URL:         "https://github.com/marcos-nsantos/test",
		UserID:      userEntity.ID,
	}
	repo := NewRepo(db)
	err = repo.Insert(context.Background(), project)
	assert.NoError(t, err)
	assert.NotEmpty(t, project.ID)
	assert.NotEmpty(t, project.CreatedAt)
	assert.NotEmpty(t, project.UpdatedAt)
}

func TestFindAll(t *testing.T) {
	repo := NewRepo(db)
	projects, err := repo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, projects)
}

func TestFindByID(t *testing.T) {
	repo := NewRepo(db)
	project, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), project.ID)
	assert.Equal(t, "Test", project.Name)
	assert.Equal(t, "Test", project.Description)
	assert.Equal(t, "https://github.com/marcos-nsantos/test", project.URL)
}

func TestUpdate(t *testing.T) {
	repo := NewRepo(db)
	project, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	project.Name = "Test Updated"
	err = repo.Update(context.Background(), project)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	repo := NewRepo(db)
	err := repo.Delete(context.Background(), 1)
	assert.NoError(t, err)
}
