package presenter

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONInternalServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("An error occurred internally"))
}

func JSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			JSONInternalServerError(w, err)
		}
	}
}

func JSONErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	JSONResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

func JSONValidationResponse(w http.ResponseWriter, fields map[string][]string) {
	response := make(map[string]any)
	response["errors"] = fields
	JSONResponse(w, http.StatusBadRequest, response)
}
