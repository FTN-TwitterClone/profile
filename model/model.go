package model

import "time"

//Info from JWT token
type AuthUser struct {
	Username string
	Role     string
	Exp      time.Time
}

//User register form
type RegisterUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Town      string `json:"town"`
	Gender    string `json:"gender"`
}

//Business user register form
type RegisterBusinessUser struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Website     string `json:"website"`
	CompanyName string `json:"companyName"`
}
