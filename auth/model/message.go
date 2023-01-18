package model

type PingResponse struct {
	Message string `json:"message"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type StandardError struct {
	Error string `json:"error"`
}

type UnauthorizedError struct {
	Error string `json:"error"`
}

type InternalServerError struct {
	Error string `json:"error"`
}
