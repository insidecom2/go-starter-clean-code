package usecase

import (
    "context"
    "errors"

    "github.com/google/uuid"
    "github.com/yourname/go-starter/scaffold/internal/user/repo"
    "golang.org/x/crypto/bcrypt"
)

var ErrUserExists = errors.New("user already exists")

// Usecase encapsulates user business logic
type Usecase struct{
    repo *repo.Repository
}

func NewUsecase(r *repo.Repository) *Usecase { return &Usecase{repo: r} }

func (u *Usecase) CreateUser(ctx context.Context, email, password string) (*repo.User, error) {
    if _, err := u.repo.GetByEmail(ctx, email); err == nil {
        return nil, ErrUserExists
    }
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil { return nil, err }
    user := &repo.User{ Email: email, PasswordHash: string(hashed) }
    if err := u.repo.Create(ctx, user); err != nil { return nil, err }
    return user, nil
}

func (u *Usecase) Authenticate(ctx context.Context, email, password string) (*repo.User, error) {
    user, err := u.repo.GetByEmail(ctx, email)
    if err != nil { return nil, err }
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return nil, err
    }
    return user, nil
}

func (u *Usecase) GetUser(ctx context.Context, id uuid.UUID) (*repo.User, error) { return u.repo.GetByID(ctx, id) }

