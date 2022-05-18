package domain

type User struct {
	Id                   string `json:"-" bson:"_id"`
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	LastVerificationCode int    `json:"-"`
	Verified             bool   `json:"-"`
	CurrentSession       string `json:"-"`
	CurrentServer        string `json:"-"`
}
