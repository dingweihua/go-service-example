package users

import "time"

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Mobile string `json:"mobile"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
}

type Users struct {
	Users []User `json:"users"`
}

type UserCreateUpdate struct {
	Username string `json:"username" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}
