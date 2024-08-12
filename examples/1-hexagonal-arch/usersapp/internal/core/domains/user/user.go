package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	UserName Username
}

func NewUser(un Username) *User {
	return &User{
		ID:       uuid.New(),
		UserName: un,
	}
}
