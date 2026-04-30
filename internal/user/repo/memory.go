package repo

import (
	"sync"

	"github.com/example/go-starter/internal/user"
)

// MemoryRepo is a threadsafe in-memory repo for development/testing
type MemoryRepo struct {
	mu    sync.RWMutex
	users map[string]user.User
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{users: make(map[string]user.User)}
}

func (r *MemoryRepo) GetAll() []user.User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := make([]user.User, 0, len(r.users))
	for _, u := range r.users {
		res = append(res, u)
	}
	return res
}

func (r *MemoryRepo) Save(u user.User) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[u.ID] = u
}
