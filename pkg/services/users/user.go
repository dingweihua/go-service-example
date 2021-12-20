package users

import (
	"context"
	"github.com/kott/go-service-example/pkg/utils/log"
)

type Repo interface {
	Get(ctx context.Context, id string) (User, error)
	GetAll(ctx context.Context, limit, offset int) ([]User, error)
	Create(ctx context.Context, user UserCreateUpdate) (string, error)
	Update(ctx context.Context, update UserCreateUpdate, id string) error
}

type Service interface {
	Get(ctx context.Context, id string) (User, error)
	GetAll(ctx context.Context, limit, offset int) ([]User, error)
	Create(ctx context.Context, user UserCreateUpdate) (User, error)
	Update(ctx context.Context, user UserCreateUpdate, id string) (User, error)
}

type UserService struct {
	repo Repo
}

func New(repo Repo) Service {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Get(ctx context.Context, id string) (User, error) {
	return s.repo.Get(ctx, id)
}

func (s *UserService) GetAll(ctx context.Context, limit, offset int) ([]User, error) {
	return s.repo.GetAll(ctx, limit, offset)
}

func (s *UserService) Create(ctx context.Context, user UserCreateUpdate) (User, error) {
	var u User
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		log.Error(ctx, "Create user service error %s", err.Error())
		return u, err
	}
	return s.repo.Get(ctx, id)
}

func (s *UserService) Update(ctx context.Context, user UserCreateUpdate, id string) (User, error) {
	var u User
	err := s.repo.Update(ctx, user, id)
	if err != nil {
		log.Error(ctx, "Update user service error %s", err.Error())
		return u, err
	}
	return s.repo.Get(ctx, id)
}
