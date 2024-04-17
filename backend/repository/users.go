package repository

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"time"
	"user-management-api/domain"
)

type UsersRepository interface {
	Create(user *domain.User) error                                                  // Adds a new user to the repository.
	GetAll(page int, pageSize int) ([]*domain.User, error)                           // Retrieves all users with pagination.
	GetByID(id uint) (*domain.User, error)                                           // Retrieves a user by their ID.
	Update(user *domain.User) error                                                  // Updates an existing user.
	HardDelete(id uint) error                                                        // Deletes the user and removes from table
	Delete(id uint) error                                                            // Soft deletes a user by marking them as deleted.
	SearchUsers(query string, page int, pageSize int) ([]*domain.User, int64, error) // Searches users based on a query string with pagination.
	Count() (int64, error)                                                           // Counts the total number of active users.
}

type users struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UsersRepository {
	return &users{db}
}

func isDuplicateEmailError(err error) bool {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) && pgError.Code == "23505" {
		return true
	}
	return false
}

func (ur *users) Create(user *domain.User) error {
	// Attempt to create the user in the database
	result := ur.db.Select("name", "age", "email").Create(user)
	if result.Error != nil {
		// Check if the error is due to a duplicate email
		if isDuplicateEmailError(result.Error) {
			return errors.New("email already exists")
		}
		return result.Error
	}
	return nil
}

func (ur *users) GetAll(page int, pageSize int) ([]*domain.User, error) {
	var users []*domain.User
	query := ur.db.Where("deleted IS NULL")
	addPagination(query, page, pageSize)
	result := query.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur *users) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	result := ur.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	if user.Deleted != 0 {
		return nil, errors.New("user is deleted")
	}
	return &user, nil
}

func (ur *users) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := ur.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	if user.Deleted != 0 {
		return nil, errors.New("user is deleted")
	}
	return &user, nil
}

func (ur *users) Update(user *domain.User) error {
	// Check if the user exists and is not deleted
	_, err := ur.GetByID(user.ID)
	if err != nil {
		return err // Handle the error (e.g., user not found)
	}

	// Update the user
	result := ur.db.Model(user).Select("name", "age", "email").Updates(user)
	if result.Error != nil {
		// Check if the error is due to a duplicate email
		if isDuplicateEmailError(result.Error) {
			return errors.New("email already exists")
		}
		return result.Error
	}
	return nil

}

func (ur *users) Count() (int64, error) {
	var count int64
	result := ur.db.Model(domain.User{}).Where("deleted IS NULL").Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (ur *users) HardDelete(id uint) error {
	result := ur.db.Delete(&domain.User{ID: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *users) Delete(id uint) error {
	user := domain.User{ID: id, Deleted: time.Now().Unix()}
	result := ur.db.Model(user).Select("deleted").Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func addPagination(query *gorm.DB, page int, pageSize int) {
	if page > 0 && pageSize > 0 {
		// Calculate the offset based on page number and page size
		offset := (page - 1) * pageSize

		// Apply pagination and fetch the subset of results
		query = query.Offset(offset).Limit(pageSize)
	}
}

func (ur *users) SearchUsers(search string, page int, pageSize int) ([]*domain.User, int64, error) {
	var users []*domain.User
	query := ur.db.Model(domain.User{}).Where("deleted is null AND (name LIKE ? OR email LIKE ? OR CAST(age AS TEXT) LIKE ?)",
		"%"+search+"%", "%"+search+"%", "%"+search+"%")
	var count int64
	query.Count(&count)
	//// Add search criteria for each column
	//query = query.Or("name LIKE ?", "%"+search+"%")
	//query = query.Or("email LIKE ?", "%"+search+"%")
	//query = query.Or("CAST(age AS TEXT) LIKE ?", "%"+search+"%") // Convert age to text for searching

	addPagination(query, page, pageSize)
	result := query.Find(&users)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	return users, count, nil
}
