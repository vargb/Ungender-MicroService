package postgres

import (
	"fmt"
	"hopeugetknowuwont/models"

	"github.com/go-pg/pg/v9"
)

type UsersRepo struct {
	DB *pg.DB
}

func (u *UsersRepo) GetUserbyID(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v=?", id), user.UserID).First()
	return &user, err
}

func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
