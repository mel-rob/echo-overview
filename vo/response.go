package vo

//Response
type HealthCheckResponse struct {
	Message string `json:"message"`
}

type Gopher struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
