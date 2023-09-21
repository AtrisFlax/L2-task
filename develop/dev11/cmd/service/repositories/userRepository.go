package repositories

import (
	"github.com/google/uuid"
	"my_httpServer/cmd/service/entities"
	"sync"
)

type UserRepository struct {
	users map[uuid.UUID]entities.UserEvents
	sync.Mutex
}

func NewUserRepository() *UserRepository {
	users := make(map[uuid.UUID]entities.UserEvents)
	return &UserRepository{users: users}
}

func (ur *UserRepository) CreateUser(newUser entities.UserEvents) {
	ur.Lock()
	defer ur.Unlock()

	ur.users[newUser.UserID] = newUser
}

func (ur *UserRepository) DeleteUser(userID uuid.UUID) {
	ur.Lock()
	defer ur.Unlock()

	delete(ur.users, userID)
}

func (ur *UserRepository) GetUser(userID uuid.UUID) entities.UserEvents {
	ur.Lock()
	defer ur.Unlock()

	return ur.users[userID]
}
