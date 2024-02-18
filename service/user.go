package service

import (
	"fmt"
	mt "gofiber-sqlx/model/user"
	"gofiber-sqlx/repository"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserService is the interface that provides user-related methods.
type User interface {
	CreateUser(name, email string) (*mt.User, error)
	GetUserByID(id int) (*mt.User, error)
	UpdateUser(id int, name, email string) error
	DeleteUser(id int) error
	GetUsers(offset, limit int) ([]mt.User, error)
}

// UserServiceImpl is the implementation of UserService.
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new UserService instance.
func NewUserService(userRepo repository.UserRepository) User {
	return &UserService{userRepo: userRepo}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(name, email string) (*mt.User, error) {
	user := &mt.User{Name: name, Email: email}

	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return &mt.User{
		ID:    id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

// GetUserByID retrieves a user by ID.
func (s *UserService) GetUserByID(id int) (*mt.User, error) {
	return s.userRepo.GetUserByID(id)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(id int, name, email string) error {
	user, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	user.Name = name
	user.Email = email

	return s.userRepo.UpdateUser(user)
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(id int) error {
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}

	return s.userRepo.DeleteUser(id)
}

func (s *UserService) GetUsers(offset, limit int) ([]mt.User, error) {
	errCh := make(chan error)
	resCh := make(chan []mt.User)
	var result []mt.User

	go func() {
		users, err := s.userRepo.GetUsersRepo(offset, limit)
		if err != nil {
			errCh <- err
			resCh <- nil
			return
		}

		resCh <- users
	}()

	// Use select to wait for either an error or a result
	select {
	case err := <-errCh:
		res := <-resCh
		logrus.WithFields(logrus.Fields{
			"debug":   fmt.Sprintf("%s", res),
			"service": "CreateKliringTransaction",
		}).WithError(err).Errorf("Update Task Data Failed")
		return nil, status.Errorf(codes.Internal, "Parse Task To Payload Failed")

	case res := <-resCh:
		result = res
		break
	}

	fmt.Println("result: ", result)
	return result, nil

}
