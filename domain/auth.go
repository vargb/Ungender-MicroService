package domain

import (
	"context"
	"errors"
	"hopeugetknowuwont/models"

	"github.com/sirupsen/logrus"
)

func (d *Domain) Login(ctx context.Context, input models.Login) (*models.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserbyID(input.UserID)
	logrus.Info("Getting User details")
	if err != nil {
		logrus.Error("Error in getting User details", err)
		return nil, ErrBadCredentials
	}

	logrus.Info("Comparing Passwords")
	err = user.ComparePassword(input.Password)
	if err != nil {
		logrus.Error("Error in comparing passwords", err)
		return nil, ErrBadCredentials
	}

	logrus.Info("Generating Token...")
	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (d *Domain) Register(ctx context.Context, input models.Signup) (*models.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserbyID(input.UserID)
	if err == nil {
		logrus.Error("userID already in use")
		return nil, errors.New("userID already in used")
	}

	user := &models.User{
		UserID: input.UserID,
		Fname:  input.Fname,
		Lname:  input.Lname,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		logrus.Error("error while hashing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		logrus.Error("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}
	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		logrus.Error("error creating a user: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		logrus.Error("error while commiting: %v", err)
		return nil, err
	}

	logrus.Info("Generating Signup Token...")
	token, err := user.GenToken()
	if err != nil {
		logrus.Error("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
