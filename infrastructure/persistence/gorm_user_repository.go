package persistence

import (
	"crud/config"
	userdomain "crud/domain/user"
	"crud/models"
)

type GormUserRepository struct{}

func NewGormUserRepository() *GormUserRepository {
	return &GormUserRepository{}
}

func (r *GormUserRepository) Save(u *userdomain.User) error {
	var m models.User
	if u.ID != 0 {
		if err := config.DB.First(&m, u.ID).Error; err != nil {
			return err
		}
	}
	m.Code = u.Code
	m.Username = u.Username
	m.PasswordHash = u.Password
	m.Email = u.Email
	m.CreatedBy = u.CreatedBy
	m.UpdatedBy = u.UpdatedBy
	m.CreatedAt = u.CreatedAt
	m.UpdatedAt = u.UpdatedAt
	m.Status = u.Status

	if u.ID == 0 {
		return config.DB.Create(&m).Error
	}
	return config.DB.Save(&m).Error
}

func (r *GormUserRepository) FindByID(id uint) (*userdomain.User, error) {
	var m models.User
	if err := config.DB.First(&m, id).Error; err != nil {
		return nil, err
	}
	return &userdomain.User{
		ID:        m.ID,
		Code:      m.Code,
		Username:  m.Username,
		Password:  m.PasswordHash,
		Email:     m.Email,
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Status:    m.Status,
	}, nil
}

func (r *GormUserRepository) Delete(id uint) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormUserRepository) CountByUsernameOrEmail(username, email string) (int64, error) {
	var count int64
	if err := config.DB.Model(&models.User{}).Where("username = ?", username).Or("email = ?", email).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *GormUserRepository) List(offset, limit int) ([]userdomain.User, int64, error) {
	var ms []models.User
	var total int64
	if err := config.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := config.DB.Limit(limit).Offset(offset).Find(&ms).Error; err != nil {
		return nil, 0, err
	}
	users := make([]userdomain.User, 0, len(ms))
	for _, m := range ms {
		users = append(users, userdomain.User{
			ID:        m.ID,
			Code:      m.Code,
			Username:  m.Username,
			Password:  m.PasswordHash,
			Email:     m.Email,
			CreatedBy: m.CreatedBy,
			UpdatedBy: m.UpdatedBy,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
			Status:    m.Status,
		})
	}
	return users, total, nil
}
