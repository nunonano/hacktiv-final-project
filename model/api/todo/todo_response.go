package api

import "time"

type TodoResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Complete  bool      `json:"complete"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
