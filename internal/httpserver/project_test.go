//go:build e2e

package httpserver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/marcos-nsantos/portfolio-api/internal/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB

func executeRequest(req *http.Request, s *Server) *httptest.ResponseRecorder {
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equalf(t, expected, actual, "Expected response code %d. Got %d\n", expected, actual)
}

func TestMain(t *testing.M) {
	client, _ := database.New()
	client.CreateTables()
	db = client.Client
	code := t.Run()
	os.Exit(code)
}

func TestCreate(t *testing.T) {
	s := CreateNewServer(db)
	s.MountHandlers()

	body := []byte(`{"name":"test","description":"test","url":"https://github.com/marcos-nsantos/portfolio-api"}`)
	req, err := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(body))
	assert.NoError(t, err)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusCreated, response.Code)
}

func TestGetProject(t *testing.T) {
	s := CreateNewServer(db)
	s.MountHandlers()

	req, err := http.NewRequest(http.MethodGet, "/projects/1", nil)
	assert.NoError(t, err)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetProjects(t *testing.T) {
	s := CreateNewServer(db)
	s.MountHandlers()

	req, err := http.NewRequest(http.MethodGet, "/projects", nil)
	assert.NoError(t, err)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateProject(t *testing.T) {
	s := CreateNewServer(db)
	s.MountHandlers()

	body := []byte(`{"name":"test","description":"test","url":"https://github.com/marcos-nsantos/portfolio-api-rest"}`)
	req, err := http.NewRequest(http.MethodPut, "/projects/1", bytes.NewBuffer(body))
	assert.NoError(t, err)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteProject(t *testing.T) {
	s := CreateNewServer(db)
	s.MountHandlers()

	req, err := http.NewRequest(http.MethodDelete, "/projects/1", nil)
	assert.NoError(t, err)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusNoContent, response.Code)
}
