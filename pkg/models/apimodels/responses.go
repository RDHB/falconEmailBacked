package apimodels

// ResponseDescriptionOne models in order to response when calling an api
type ResponseDescriptionOne struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	TotalData   int64  `json:"total_data"`
	MaxDataPage int64  `json:"max_data_page"`
	TotalPages  int64  `json:"total_pages"`
	Page        int64  `json:"page"`
	Data        any    `json:"data"`
}
