package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
)

// UserRepository interface
type UserRepository interface {
	All(context.Context) ([]*entity.User, error)
	Create(context.Context, *entity.User) (*entity.User, error)
	FindByID(context.Context, uuid.UUID) (*entity.User, error)
	FindByEmail(context.Context, string) (*entity.User, error)
	ChangePassword(context.Context, uuid.UUID, string) (*entity.User, error)
}

// New func
func New() UserRepository {
	db := ioc.Get(database.PostgreSQL{}).DB

	return newSQL(db)
}
