package repo

import (
    "context"
    "time"

    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
)

// User represents DB model for users
type User struct {
    ID           uuid.UUID `db:"id"`
    Email        string    `db:"email"`
    PasswordHash string    `db:"password_hash"`
    CreatedAt    time.Time `db:"created_at"`
}

// Repository provides access to user storage
type Repository struct{
    db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Create(ctx context.Context, u *User) error {
    if u.ID == uuid.Nil {
        u.ID = uuid.New()
    }
    _, err := r.db.NamedExecContext(ctx, `INSERT INTO users (id, email, password_hash, created_at) VALUES (:id, :email, :password_hash, :created_at)`, u)
    return err
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
    var u User
    if err := r.db.GetContext(ctx, &u, "SELECT id, email, password_hash, created_at FROM users WHERE id=$1", id); err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
    var u User
    if err := r.db.GetContext(ctx, &u, "SELECT id, email, password_hash, created_at FROM users WHERE email=$1", email); err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
    _, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
    return err
}

func (r *Repository) UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) error {
    _, err := r.db.ExecContext(ctx, "UPDATE users SET password_hash=$2 WHERE id=$1", id, passwordHash)
    return err
}
