package users

import (
	"errors"
	"regexp"
	"user-management-api/domain"
	"user-management-api/repository"
)

type Service interface {
	CreateUser(user *domain.User) error
	GetUsers(page int, pageSize int) ([]*domain.User, error)
	GetUser(userId int) (*domain.User, error)
	UpdateUser(userId int, user *domain.User) error
	DeleteUser(userId int) error
	SearchUsers(query string, page int, pageSize int) ([]*domain.User, int64, error)
	Count() (int64, error)
}

type service struct {
	users repository.UsersRepository
}

func NewService(users repository.UsersRepository) Service {
	return &service{
		users: users,
	}
}

func (s *service) CreateUser(user *domain.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}
	// Additional business logic/validation before creating user, if needed

	if err := s.users.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *service) GetUsers(page int, pageSize int) ([]*domain.User, error) {
	return s.users.GetAll(page, pageSize)
}

func (s *service) GetUser(userId int) (*domain.User, error) {
	return s.users.GetByID(uint(userId))
}

func (s *service) UpdateUser(userId int, user *domain.User) error {
	err := s.validateUser(user)
	if err != nil {
		return err
	}
	user.ID = uint(userId)
	err = s.users.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteUser(userId int) error {
	err := s.users.Delete(uint(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s *service) SearchUsers(query string, page int, pageSize int) ([]*domain.User, int64, error) {
	return s.users.SearchUsers(query, page, pageSize)
}

func (s *service) Count() (int64, error) {
	return s.users.Count()
}

func (s *service) validateUser(user *domain.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	} else if !isValidEmail(user.Email) {
		return errors.New("invalid email")
	}
	if user.Age < 18 {
		return errors.New("age must be 18 or older")
	}
	// Additional validation rules can be added here
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}
