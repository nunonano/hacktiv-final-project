package api

type TodoRequest struct {
	Name     string `json:"name"`
	Complete string `json:"complete"`
}
