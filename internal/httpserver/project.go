package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/marcos-nsantos/portfolio-api/internal/entity"
	"github.com/marcos-nsantos/portfolio-api/internal/httpserver/presenter"
	"github.com/marcos-nsantos/portfolio-api/internal/validator"
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
		presenter.JSONErrorResponse(w, http.StatusBadRequest, err)
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
