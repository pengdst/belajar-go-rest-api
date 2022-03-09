package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/db/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/util/di"
)

// UserService interface
type UserService interface {
	All(context.Context) ([]*entity.User, error)
	Create(context.Context, *entity.User) (*entity.User, error)
	FindByID(context.Context, uuid.UUID) (*entity.User, error)
	FindByEmail(context.Context, string) (*entity.User, error)
	ChangePassword(context.Context, uuid.UUID, string) (*entity.User, error)
}

// userService struct
type userService struct {
	repo repository.Repository
}

// NewUserService func
func NewUserService(ioc di.IOC) UserService {
	repo := getRepository(ioc)
	return userService{
		repo,
	}
}

// All func
func (u userService) All(c context.Context) (users []*entity.User, err error) {
	usersFromTable, err := u.repo.User.All(c)
	if err != nil {
		return nil, err
	}

	for _, uft := range usersFromTable {
		users = append(users, entity.UserFromTable(uft))
	}

	return users, nil
}

// Create func
func (u userService) Create(c context.Context, user *entity.User) (*entity.User, error) {
	userFromTable, err := u.repo.User.Create(c, user.ToTable())
	if err != nil {
		return nil, err
	}

	return entity.UserFromTable(userFromTable), nil
}

// FindByEmail func
func (u userService) FindByEmail(c context.Context, email string) (*entity.User, error) {
	userFromTable, err := u.repo.User.FindByEmail(c, email)
	if err != nil {
		return nil, err
	}

	return entity.UserFromTable(userFromTable), nil
}

// FindByID func
func (u userService) FindByID(c context.Context, id uuid.UUID) (*entity.User, error) {
	userFromTable, err := u.repo.User.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	return entity.UserFromTable(userFromTable), nil
}

// ChangePassword func
func (u userService) ChangePassword(c context.Context, id uuid.UUID, password string) (*entity.User, error) {
	userFromTable, err := u.repo.User.ChangePassword(c, id, password)
	if err != nil {
		return nil, err
	}

	return entity.UserFromTable(userFromTable), nil
}
