package repositories

import (
	"github.com/najayi/learning-go-shop/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetByID(id uint) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetByEmailAndActive(email string, isActive bool) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Create(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Update(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) CreateRefreshToken(token *models.RefreshToken) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) GetValidRefreshToken(token string) (*models.RefreshToken, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) DeleteRefreshToken(token string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) DeleteRefreshTokenByID(id uint) error {
	//TODO implement me
	panic("implement me")
}
