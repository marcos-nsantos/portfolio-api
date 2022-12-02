//go:build e2e

package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcos-nsantos/portfolio-api/internal/database"
	"github.com/marcos-nsantos/portfolio-api/internal/httpserver"
	"github.com/stretchr/testify/assert"
)

var s *httpserver.Server

func executeRequest(req *http.Request, s *httpserver.Server) *httptest.ResponseRecorder {
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equalf(t, expected, actual, "Expected response code %d. Got %d\n", expected, actual)
}

func TestMain(t *testing.M) {
	db, _ := database.New()
	db.CreateTables()
	s = httpserver.CreateNewServer(db.Client)
	s.MountHandlers()
}

func TestCreate(t *testing.T) {
	body := []byte(`{"name":"test","description":"test","url":"https://github.com/marcos-nsantos/portfolio-api"}`)
	req, err := http.NewRequest(http.MethodPost, "/projects", bytes.NewBuffer(body))
	assert.NoError(t, err)
	response := executeRequest(req, s)
	checkResponseCode(t, http.StatusCreated, response.Code)
}
