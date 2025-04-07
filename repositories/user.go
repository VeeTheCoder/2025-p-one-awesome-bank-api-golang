package repositories

import (
    "github-api-boilerplate/models"
    "google.golang.org/protobuf/proto"
)

type UserRepository struct {
    db *gorm.DB
}

func (u *UserRepository) Create(user models.User) (*models.User, error) {
    data, err := proto.Marshal(&user)
    if err != nil {
        return nil, err
    }
    return u.db.Create(&models.User{Data: data}).Find(&models.User{}).Error
}

func (u *UserRepository) GetAll() ([]*models.User, error) {
    var users []*models.User
    err := u.db.Find(&users).Error
    return users, err
}

func (u *UserRepository) Update(user models.User) (*models.User, error) {
    data, err := proto.Marshal(&user)
    if err != nil {
        return nil, err
    }
    return u.db.Model(&models.User{}).
        Update("data", data).
        Find(&models.User{}).Error
}

func (u *UserRepository) Delete(user models.User) error {
    return u.db.Delete(&models.User{}).Error
}