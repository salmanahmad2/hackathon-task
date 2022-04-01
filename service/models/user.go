package models

type User struct {
	UserId   *string `json:"user_id" db:"user_id"`
	Email    *string `json:"email" db:"email"`
	Password *string `json:"password" db:"password"`
}

func NewUser() *User {
	return &User{
		UserId:   new(string),
		Email:    new(string),
		Password: new(string),
	}
}

type Todo struct {
	TodoId      *string `json:"todo_id" db:"todo_id"`
	UserId      *string `json:"user_id" db:"user_id"`
	Title       *string `json:"title" db:"title"`
	Discription *string `json:"discription" db:"discription"`
	Status      *string `json:"status" db:"status"`
}

func NewTodo() *Todo {
	return &Todo{
		UserId:      new(string),
		TodoId:      new(string),
		Title:       new(string),
		Discription: new(string),
		Status:      new(string),
	}
}

type SignUpRequest struct {
	Email    *string `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"required,password"`
	OTP      *string `json:"otp"`
}

func NewSignUpRequest() *SignUpRequest {
	return &SignUpRequest{
		Email:    new(string),
		Password: new(string),
		OTP:      new(string),
	}
}

type SignInRequest struct {
	Email    *string `json:"email" db:"email" validate:"required,gt=0"`
	Password *string `json:"password" db:"password" validate:"required,password"`
}

func NewSignInRequest() *SignInRequest {
	return &SignInRequest{
		Email:    new(string),
		Password: new(string),
	}
}
