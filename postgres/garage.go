package postgres

import (
	"hopeugetknowuwont/models"

	"github.com/go-pg/pg/v9"
	"github.com/sirupsen/logrus"
)

type GarageRepo struct {
	DB *pg.DB
}

var userRep *UsersRepo

func (g *GarageRepo) GetbyID(id string) (*models.Garage, error) {
	var gar models.Garage
	err := g.DB.Model(&gar).Where("carid = ?", id).First()
	return &gar, err
}

func (g *GarageRepo) GetCartoUser(getcar *models.GetCar) error {
	us, err := userRep.GetUserbyID(getcar.UserID)
	if err != nil {
		logrus.Error("error in getting userid", err)
		return err
	}
	_, err = userRep.DB.Model(us).Where("carid = ?", getcar.CarID).Update()
	if err != nil {
		logrus.Error("error in adding carid to user", err)
		return err
	}
	gar, err := g.GetbyID(getcar.CarID)
	if err != nil {
		return err
	}
	_, err = g.DB.Model(gar).Where("available = ?", false).Update()
	return err
}

func (g *GarageRepo) DelCarfromUser(delcar *models.GetCar) error {
	us, err := userRep.GetUserbyID(delcar.UserID)
	if err != nil {
		logrus.Error("error in getting userid", err)
		return err
	}
	_, err = userRep.DB.Model(us).Where("carid = ?", delcar.CarID).Delete()
	if err != nil {
		logrus.Error("error in deleting carid to user", err)
		return err
	}
	gar, err := g.GetbyID(delcar.CarID)
	if err != nil {
		return err
	}
	_, err = g.DB.Model(gar).Where("available = ?", true).Update()
	return err
}

func (g *GarageRepo) GetAvalCars(user *models.User) ([]*models.Garage, error) {
	var gar []*models.Garage
	err := g.DB.Model(&gar).Where("available = ?", true).Order("carid").Select()
	return gar, err
}
