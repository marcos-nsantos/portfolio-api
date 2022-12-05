package user

import (
	"context"
	"fmt"
	"github.com/marcos-nsantos/portfolio-api/internal/database"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
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
	user := &entity.User{
		FirstName: "Marcos",
		LastName:  "Santos",
		Email:     "email@email.com",
		Password:  "password",
	}
	repo := NewRepo(db)
	err := repo.Insert(context.Background(), user)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.ID)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
}

func TestFindAll(t *testing.T) {
	repo := NewRepo(db)
	users, err := repo.FindAll(context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, users)
}

func TestFindByID(t *testing.T) {
	repo := NewRepo(db)
	user, err := repo.FindByID(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), user.ID)
	assert.Equal(t, "Marcos", user.FirstName)
	assert.Equal(t, "Santos", user.LastName)
	assert.Equal(t, "email@email.com", user.Email)
}

func TestUpdate(t *testing.T) {
	user := &entity.User{
		FirstName: "Marcos",
		LastName:  "Santos",
		Email:     "email@email.com",
	}
	repo := NewRepo(db)
	err := repo.Insert(context.Background(), user)
	assert.NoError(t, err)

	user.Email = "user@email.com"
	err = repo.Update(context.Background(), user)
	assert.NoError(t, err)

	user, err = repo.FindByID(context.Background(), user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "user@email.com", user.Email)
}
