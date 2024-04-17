package users

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"user-management-api/domain"
)

// MockUsersRepository is a mock type for the UsersRepository
type MockUsersRepository struct {
	mock.Mock
}

// Create mocks the repository method
func (m *MockUsersRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// GetAll mocks the repository method
func (m *MockUsersRepository) GetAll(page int, pageSize int) ([]*domain.User, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*domain.User), args.Error(1)
}

// GetByID mocks the repository method
func (m *MockUsersRepository) GetByID(id uint) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

// Update mocks the repository method
func (m *MockUsersRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// HardDelete mocks the repository method
func (m *MockUsersRepository) HardDelete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Delete mocks the repository method
func (m *MockUsersRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// SearchUsers mocks the repository method
func (m *MockUsersRepository) SearchUsers(query string, page int, pageSize int) ([]*domain.User, int64, error) {
	args := m.Called(query, page, pageSize)
	return args.Get(0).([]*domain.User), args.Get(1).(int64), args.Error(2)
}

// Count mocks the repository method
func (m *MockUsersRepository) Count() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	svc := NewService(mockRepo)
	validUser := &domain.User{Name: "John Doe", Email: "john@example.com", Age: 30}
	invalidEmailUser := &domain.User{Name: "John Doe", Email: "bademail", Age: 30}
	youngUser := &domain.User{Name: "John Doe", Email: "john@example.com", Age: 17}

	testCases := []struct {
		name    string
		user    *domain.User
		mock    func()
		wantErr bool
		errMsg  string
	}{
		{
			name: "Successful Create",
			user: validUser,
			mock: func() {
				mockRepo.On("Create", validUser).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "Invalid Email",
			user:    invalidEmailUser,
			mock:    func() {},
			wantErr: true,
			errMsg:  "invalid email",
		},
		{
			name:    "Age Below 18",
			user:    youngUser,
			mock:    func() {},
			wantErr: true,
			errMsg:  "age must be 18 or older",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := svc.CreateUser(tc.user)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockUsersRepository)
	svc := NewService(mockRepo)
	validUser := &domain.User{ID: 1, Name: "Jane Doe", Email: "jane@example.com", Age: 25}
	invalidUser := &domain.User{ID: 1, Name: "", Email: "jane@example.com", Age: 25}

	testCases := []struct {
		name    string
		userId  int
		user    *domain.User
		mock    func()
		wantErr bool
		errMsg  string
	}{
		{
			name:   "Successful Update",
			userId: 1,
			user:   validUser,
			mock: func() {
				mockRepo.On("Update", validUser).Return(nil)
			},
			wantErr: false,
		},
		{
			name:    "Invalid User Data",
			userId:  1,
			user:    invalidUser,
			mock:    func() {},
			wantErr: true,
			errMsg:  "name is required",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			err := svc.UpdateUser(tc.userId, tc.user)
			if tc.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
