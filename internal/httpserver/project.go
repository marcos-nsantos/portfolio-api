package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/marcos-nsantos/portfolio-api/internal/errs"
	"github.com/marcos-nsantos/portfolio-api/internal/httpserver/presenter"
	"github.com/marcos-nsantos/portfolio-api/internal/validator"
	"gorm.io/gorm"
)

type ProjectRequest struct {
	Name        string `json:"name" validate:"required,notblank,max=100"`
	Description string `json:"description" validate:"required,notblank,max=255"`
	URL         string `json:"url" validate:"required,url,max=255"`
}

func (p *ProjectRequest) ToEntity() *entity.Project {
	return &entity.Project{
		Name:        p.Name,
		Description: p.Description,
		URL:         p.URL,
	}
}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	var request ProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidBodyRequest)
		return
	}

	if ok, errors := validator.Validate(request); !ok {
		presenter.JSONValidationResponse(w, errors)
		return
	}

	toEntity := request.ToEntity()
	if err := s.Project.Create(r.Context(), toEntity); err != nil {
		presenter.JSONInternalServerError(w, err)
		return
	}

	projectPresenter := presenter.NewProjectPresenter(toEntity)
	presenter.JSONResponse(w, http.StatusCreated, projectPresenter)
}

func (s *Server) getProject(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idUint, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidID)
		return
	}

	project, err := s.Project.GetByID(r.Context(), uint(idUint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			presenter.JSONErrorResponse(w, http.StatusNotFound, err)
			return
		}
		presenter.JSONInternalServerError(w, err)
		return
	}

	projectPresenter := presenter.NewProjectPresenter(project)
	presenter.JSONResponse(w, http.StatusOK, projectPresenter)
}
