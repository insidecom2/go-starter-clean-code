package usecase

import "github.com/example/go-starter/internal/user"

// UserRepo defines the subset of repository needed by the usecase
type UserRepo interface {
	GetAll() []user.User
	Save(user.User)
}

// UserUsecase contains business logic around users
type UserUsecase struct {
	repo UserRepo
}

func New(r UserRepo) *UserUsecase { return &UserUsecase{repo: r} }

func (uc *UserUsecase) List() []user.User { return uc.repo.GetAll() }

func (uc *UserUsecase) Create(u user.User) { uc.repo.Save(u) }
