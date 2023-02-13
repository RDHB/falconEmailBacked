package functions

import (
	"encoding/json"
	"net/http"

	apiModels "falconEmailBackend/pkg/models/apimodels"
)

// WriteResponseOne funcition used for write errors within api
func WriteResponseOne(w http.ResponseWriter, code int, message string, totalData int64, maxDataPage int64, totalPages int64, page int64, data any) {
	responseJSON := apiModels.ResponseDescriptionOne{
		Code:        code,
		Message:     message,
		TotalData:   totalData,
		MaxDataPage: maxDataPage,
		TotalPages:  totalPages,
		Page:        page,
		Data:        data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseJSON)
}
