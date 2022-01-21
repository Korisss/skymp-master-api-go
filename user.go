package master_api

type User struct {
	Id             int    `json:"-" db:"id"`
	Name           string `json:"name" binding: "required"`
	Email          string `json:"email" binding: "required"`
	Password       string `json:"password" binding: "required"`
	Verified       bool   `json:"-"`
	CurrentSession string `json:"-"`
	CurrentServer  string `json:"-"`
}
