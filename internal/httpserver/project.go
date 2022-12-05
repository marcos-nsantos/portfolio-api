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

type requestCreate struct {
	Name        string `json:"name" validate:"required,notblank,max=100"`
	Description string `json:"description" validate:"required,notblank,max=255"`
	URL         string `json:"url" validate:"required,url,max=255"`
	UserID      uint64 `json:"user_id" validate:"required"`
}

type requestUpdate struct {
	Name        string `json:"name" validate:"required,notblank,max=100"`
	Description string `json:"description" validate:"required,notblank,max=255"`
	URL         string `json:"url" validate:"required,url,max=255"`
}

func (rc *requestCreate) entity() *entity.Project {
	return &entity.Project{
		Name:        rc.Name,
		Description: rc.Description,
		URL:         rc.URL,
		UserID:      rc.UserID,
	}
}

type requests interface {
	entity() *entity.Project
}

func convertToProjectEntity(request requests) *entity.Project {
	return request.entity()
}

func (cr *requestUpdate) entity() *entity.Project {
	return &entity.Project{
		Name:        cr.Name,
		Description: cr.Description,
		URL:         cr.URL,
	}
}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	var request requestCreate
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidBodyRequest)
		return
	}

	if ok, validationErrors := validator.Validate(request); !ok {
		presenter.JSONValidationResponse(w, validationErrors)
		return
	}

	project := convertToProjectEntity(&request)
	if err := s.Project.Create(r.Context(), project); err != nil {
		presenter.JSONInternalServerError(w, err)
		return
	}

	projectPresenter := presenter.NewProjectPresenter(project)
	presenter.JSONResponse(w, http.StatusCreated, projectPresenter)
}

func (s *Server) getProject(w http.ResponseWriter, r *http.Request) {
	idParam, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidID)
		return
	}

	project, err := s.Project.GetByID(r.Context(), idParam)
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

func (s *Server) getAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := s.Project.GetAll(r.Context())
	if err != nil {
		presenter.JSONInternalServerError(w, err)
		return
	}

	projectsPresenter := presenter.NewProjectsPresenter(projects)
	presenter.JSONResponse(w, http.StatusOK, projectsPresenter)
}

func (s *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	idParam, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidID)
		return
	}

	var request requestUpdate
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidBodyRequest)
		return
	}

	if ok, validationErrors := validator.Validate(request); !ok {
		presenter.JSONValidationResponse(w, validationErrors)
		return
	}

	project := convertToProjectEntity(&request)
	project.ID = idParam
	if err := s.Project.Update(r.Context(), project); err != nil {
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

func (s *Server) deleteProject(w http.ResponseWriter, r *http.Request) {
	idParam, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		presenter.JSONErrorResponse(w, http.StatusBadRequest, errs.ErrInvalidID)
		return
	}

	if err := s.Project.Delete(r.Context(), idParam); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			presenter.JSONErrorResponse(w, http.StatusNotFound, err)
			return
		}
		presenter.JSONInternalServerError(w, err)
		return
	}

	presenter.JSONResponse(w, http.StatusNoContent, nil)
}
