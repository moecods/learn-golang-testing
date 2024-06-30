package tutorialgolangtesting

import (
	"errors"
)

type User struct {
	ID int
	Name string
}

type UserRepository interface {
	GetUser(id int) (*User, error)
}

type MockUserRepository struct {
    users map[int]*User
}

type UserService struct {
	repo UserRepository
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.GetUser(id)

	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    user, ok := m.users[id]
    if !ok {
        return nil, errors.New("user not found")
    }
    return user, nil
}