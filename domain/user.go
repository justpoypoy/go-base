package domain

import "time"

// User for get single user
type User struct {
	ID        uint       `json:"id" form:"id"`
	Fullname  string     `json:"fullname" form:"fullname"`
	Email     string     `json:"email" form:"email"`
	Password  string     `json:"password" form:"password"`
	IsActived bool       `json:"is_active" form:"is_actived"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Users for get all user
type Users []User
