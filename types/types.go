package types

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
