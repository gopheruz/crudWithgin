package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type UpdateUser struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type GetAllUsersParams struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"page"`
	Search string `json:"search"`
}
