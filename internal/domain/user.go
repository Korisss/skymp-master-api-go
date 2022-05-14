package domain

type User struct {
	Id             string `json:"-" db:"_id"`
	Name           string `json:"name" binding:"required" db:"name"`
	Email          string `json:"email" binding:"required" db:"email"`
	Password       string `json:"password" binding:"required" db:"password_hash"`
	Verified       bool   `json:"-"`
	CurrentSession string `json:"-"`
	CurrentServer  string `json:"-"`
}
