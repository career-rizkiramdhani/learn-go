package user

import "time"

// User represents the domain entity for a user.
type User struct {
	ID        uint
	Code      string
	Username  string
	Password  string // password hash
	Email     string
	CreatedBy uint
	UpdatedBy uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    string
}

// Repository defines storage operations for User in the domain.
type Repository interface {
	Save(u *User) error
	FindByID(id uint) (*User, error)
	FindByUsername(username string) (*User, error)
	Delete(id uint) error
	CountByUsernameOrEmail(username, email string) (int64, error)
	List(offset, limit int) ([]User, int64, error)
}
