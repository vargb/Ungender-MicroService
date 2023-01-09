package domain

import (
	"context"
	"errors"
	"hopeugetknowuwont/middleware"
	"hopeugetknowuwont/models"

	"github.com/sirupsen/logrus"
)

func (d *Domain) GetCarAuth(ctx context.Context, in models.GetCar) (*models.User, error) {
	logrus.Info("Authenticating for getting car")
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}
	gar, err := d.GarageRepos.GetbyID(in.CarID)
	if err != nil || gar == nil {
		return nil, errors.New("error in getting car")
	}
	if !gar.Available {
		return nil, errors.New("carid not available")
	}
	currentUser.CarID = in.CarID
	err = d.GarageRepos.GetCartoUser(&in)
	if err != nil {
		logrus.Error("error in getting car to user", err)
		return nil, err
	}
	return currentUser, nil

}

func (d *Domain) DelCarAuth(ctx context.Context, out models.GetCar) (*models.User, error) {
	logrus.Info("Authenticating for getting car")
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}
	gar, err := d.GarageRepos.GetbyID(out.CarID)
	if err != nil || gar == nil {
		logrus.Error("error in getting carid", err)
		return nil, errors.New("error in getting car")
	}
	currentUser.CarID = ""
	err = d.GarageRepos.DelCarfromUser(&out)
	if err != nil {
		logrus.Error("error in getting car to user", err)
		return nil, err
	}
	return currentUser, nil

}
