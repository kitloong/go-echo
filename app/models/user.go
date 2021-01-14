package models

import "time"

type User struct {
	ID    int       `json:"id" form:"id" query:"id"`
	Name  string    `json:"name" form:"name" query:"name"`
	Email string    `json:"email" form:"email" query:"email"`
	Now   time.Time `json:"now" form:"now" query:"now"`
}
