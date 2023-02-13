package apimodels

// ErrorDescriptionOne models that describe structrure for error with description, first version
type ErrorDescriptionOne struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}
