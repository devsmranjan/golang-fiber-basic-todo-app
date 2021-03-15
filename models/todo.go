package models

import (
	"time"
)

// Todo - todo model
type Todo struct {
	ID        *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title     *string   `json:"title"`
	Completed *bool     `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
