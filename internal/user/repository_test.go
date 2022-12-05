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
