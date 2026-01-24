package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	userdomain "crud/domain/user"
	"crud/infrastructure/persistence"
)

var repo userdomain.Repository

func init() {
	// default wiring: GORM-backed repository
	repo = persistence.NewGormUserRepository()
}

func CreateUser(username, password, email, prefix string, createdBy uint) (*userdomain.User, error) {
	// check unique username/email
	cnt, err := repo.CountByUsernameOrEmail(username, email)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("username or email already exists")
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	u := &userdomain.User{
		Code:      GenerateUserCode(prefix),
		Username:  username,
		Password:  string(pw),
		Email:     email,
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    "active",
	}

	if err := repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserByID(id uint) (*userdomain.User, error) {
	return repo.FindByID(id)
}

func UpdateUser(id uint, username, password, email, status *string, updatedBy uint) (*userdomain.User, error) {
	u, err := repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if username != nil {
		u.Username = *username
	}
	if email != nil {
		u.Email = *email
	}
	if status != nil {
		u.Status = *status
	}
	if password != nil && *password != "" {
		pw, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		u.Password = string(pw)
	}

	u.UpdatedBy = updatedBy
	u.UpdatedAt = time.Now()

	if err := repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}

func DeleteUser(id uint) error {
	return repo.Delete(id)
}

func ListUsers(page, size int) ([]userdomain.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size
	return repo.List(offset, size)
}

// GenerateUserCode is a helper kept here for compatibility with previous code
func GenerateUserCode(prefix string) string {
	// simple timestamp-based code, keep previous behavior if any
	return prefix + time.Now().Format("20060102150405")
}
