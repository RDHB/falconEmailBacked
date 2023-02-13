package functions

import (
	"encoding/json"
	"log"
	"net/http"

	apiModels "falconEmailBackend/pkg/models/apimodels"
)

// WriteErrorOne funcition used for write errors within api
func WriteErrorOne(w http.ResponseWriter, code int, message string, details string) {
	log.Println(message)
	log.Println(details)
	errorJSON := apiModels.ErrorDescriptionOne{
		Code:    code,
		Message: message,
		Details: details,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorJSON)
}
