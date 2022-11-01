package models

import "time"

type Author struct {
	Id        string     `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateAuthorModel struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UpdateAuthorModel struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
