package db

import "sync"

type User struct {
	ID       string
	Username string
	Password string
}

type InMemoryDB struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		users: make(map[string]User),
	}
}

func (db *InMemoryDB) CreateUser(user User) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.users[user.Username]; exists {
		return ErrUserExists
	}

	db.users[user.Username] = user
	return nil
}

func (db *InMemoryDB) GetUser(username string) (User, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	user, exists := db.users[username]
	return user, exists
}

var ErrUserExists = &UserError{"user already exists"}

type UserError struct {
	msg string
}

func (e *UserError) Error() string {
	return e.msg
}