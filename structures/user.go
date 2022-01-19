package master_api

type User struct {
	Id             int    `json:"-"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Verified       bool   `json:"-"`
	CurrentSession string `json:"-"`
	CurrentServer  string `json:"-"`
}
