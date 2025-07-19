package models

type UserLogin struct {
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type Payload struct {
	User string `json:"user"`
	Role string `json:"role"`
}
