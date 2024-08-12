package ports

import (
	"context"
	"usersapp/internal/core/domains/user"
)

type UserPort interface {
	Add(ctx context.Context, u user.User) error
}
