package model

import "time"

// AuthUser Info from JWT token
type AuthUser struct {
	Username string
	Role     string
	Exp      time.Time
}

// User Combined data of business and ordinary user.
// Combining is the simplest solution, since Go is statically typed.
type User struct {
	Username    string `json:"username,omitempty" bson:"username,omitempty"`
	Email       string `json:"email,omitempty" bson:"email,omitempty"`
	FirstName   string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Town        string `json:"town,omitempty" bson:"town,omitempty"`
	Gender      string `json:"gender,omitempty" bson:"gender,omitempty"`
	Website     string `json:"website,omitempty" bson:"website,omitempty"`
	CompanyName string `json:"companyName,omitempty" bson:"companyName,omitempty"`
	Private     bool   `json:"private" bson:"private"`
}
type UpdateProfile struct {
	Username string `json:"username,omitempty" bson:"username,omitempty"`
	Private  bool   `json:"private" bson:"private"`
}
