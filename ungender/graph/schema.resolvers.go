package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	models "hopeugetknowuwont/ungender/graph/model"
	potgres "hopeugetknowuwont/ungender/pgres"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input models.Signup) (*models.User, error) {
	h := potgres.GetPqHandler()
	tx := h.DB.Where("phno = ?", input.Phno).First(&potgres.User{})
	if tx.Error != logger.ErrRecordNotFound {
		logrus.Info("User already present")
		return nil, errors.New("user already present")
	}
	hash := sha256.New()
	hash.Write([]byte(input.Phno))
	hashString := hex.EncodeToString(hash.Sum(nil))
	newUser := &models.User{
		Fname:    input.Fname,
		Lname:    input.Lname,
		Phno:     &input.Phno,
		Password: &input.Password,
		Userid:   &hashString,
	}

	res := h.DB.Create(&potgres.User{
		Fname:    *input.Fname,
		Lname:    *input.Lname,
		UserId:   hashString,
		Phno:     input.Phno,
		Password: input.Password,
	})
	if res.Error != nil {
		logrus.Error("Error in posting user", res.Error)
		return nil, res.Error
	}
	return newUser, nil
}

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, input models.Login) (*string, error) {
	c, err := GinContextFromContext(ctx)
	if err != nil {
		logrus.Error("error in fetching the gin.Context")
		return nil, err
	}
	h := potgres.GetPqHandler()

	cookie, err := c.Cookie(input.Phno)
	if cookie != "" {
		logrus.Info("User has already been logged in")
		successful := "You have already been logged in"
		return &successful, nil
	}

	var expectedUser potgres.User
	tx := h.DB.Where(&potgres.User{Phno: input.Phno, Password: input.Password}).Find(&expectedUser)
	if tx.Error != logger.ErrRecordNotFound && expectedUser.Phno != "" {
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &potgres.Claims{
			Phno: input.Phno,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(expectedUser.Phno))
		if err != nil {
			logrus.Error("error in signing the token", err)
			return nil, err
		}

		c.SetCookie(input.Phno, tokenString, 120, "/", "localhost", false, true)

		successful := "You have logged in successfully"
		return &successful, nil

	}

	logrus.Info("Unauthorized user")
	return nil, errors.New("unauthorized user")
}

// Signout is the resolver for the signout field.
func (r *mutationResolver) Signout(ctx context.Context, input models.Login) (*string, error) {
	c, err := GinContextFromContext(ctx)
	if err != nil {
		logrus.Error("error in getting gin.Context")
		return nil, err
	}

	cookie, err := c.Cookie(input.Phno)
	if cookie == "" {
		logrus.Error("idk, shouldnt you login to logout")
		return nil, errors.New("idk, shouldnt you login to logout")
	}

	c.SetCookie(input.Phno, "", -1, "/", "localhost", false, true)

	successful := "You have been logged out, Thank You!"
	return &successful, nil
}

// Getcar is the resolver for the getcar field.
func (r *mutationResolver) Getcar(ctx context.Context, input models.GetCar) (*models.User, error) {
	h := potgres.GetPqHandler()
	c, err := GinContextFromContext(ctx)
	if err != nil {
		logrus.Error("error in getting gin.Context")
		return nil, err
	}

	var expectedUser potgres.User
	tx := h.DB.Where(&potgres.User{UserId: input.Userid}).Find(&expectedUser)
	if tx.Error != nil {
		logrus.Error("Error in getting user ", tx.Error)
		return nil, tx.Error
	}

	cookie, err := c.Cookie(expectedUser.Phno)
	if cookie == "" {
		logrus.Error("idk, shouldnt you login to get the car")
		return nil, errors.New("login to get the car")
	}

	var expectedCar potgres.Garage
	tx = h.DB.Where(&potgres.Garage{Carid: input.Carid}).Find(&expectedCar)
	if tx.Error != nil {
		logrus.Error("Error in getting the car ", tx.Error)
		return nil, tx.Error
	}

	if !expectedCar.Available {
		logrus.Info("current car is unavailable, pls use a another car. To get the list use getAll query")
		return nil, errors.New("current car is unavailable, pls use a another car. To get the list use getAll query")
	}

	if expectedUser.CarId != "" {
		logrus.Info("return the car and then ask for another")
		return nil, errors.New("return the car and then ask for another")
	}

	h.DB.Model(expectedCar).Where("carid = ?", expectedCar.Carid).Update("available", false)
	h.DB.Model(expectedUser).Updates(potgres.User{CarId: input.Carid})

	newUser := &models.User{Fname: &expectedUser.Fname, Lname: &expectedUser.Lname, Userid: &expectedUser.UserId, Phno: &expectedUser.Phno, Carid: &input.Carid}
	return newUser, nil
}

// Returncar is the resolver for the returncar field.
func (r *mutationResolver) Returncar(ctx context.Context, input models.GetCar) (*models.User, error) {
	h := potgres.GetPqHandler()
	c, err := GinContextFromContext(ctx)
	if err != nil {
		logrus.Error("error in getting gin.Context")
		return nil, err
	}

	var expectedUser potgres.User
	tx := h.DB.Where(&potgres.User{UserId: input.Userid}).Find(&expectedUser)
	if tx.Error != nil {
		logrus.Error("Error in getting user ", tx.Error)
		return nil, tx.Error
	}

	cookie, err := c.Cookie(expectedUser.Phno)
	if cookie == "" {
		logrus.Error("idk, shouldnt you login to get the car")
		return nil, errors.New("login to get the car")
	}

	var expectedCar potgres.Garage
	tx = h.DB.Where(&potgres.Garage{Carid: input.Carid}).Find(&expectedCar)
	if tx.Error != nil {
		logrus.Error("Error in getting the car ", tx.Error)
		return nil, tx.Error
	}

	if expectedCar.Available || expectedCar.Carid != expectedUser.CarId {
		logrus.Info("current car cant be returned, pls use a valid car ID. To get the list use getAll query")
		return nil, errors.New("current car cant be returned, pls use a valid car ID. To get the list use getAll query")
	}

	h.DB.Model(expectedUser).Update("car_id", "")
	h.DB.Model(expectedCar).Updates(potgres.Garage{Available: true, LastUsedDate: time.Now().String()})

	newUser := &models.User{Fname: &expectedUser.Fname, Lname: &expectedUser.Lname, Userid: &expectedUser.UserId, Phno: &expectedUser.Phno, Carid: nil}
	return newUser, nil
}

// GetAll is the resolver for the getAll field.
func (r *queryResolver) GetAll(ctx context.Context) ([]*models.Garage, error) {
	var cars []*models.Garage
	h := potgres.GetPqHandler()
	if res := h.DB.Find(&cars); res.Error != nil {
		logrus.Error("Error in getting the cars", res.Error)
		return nil, res.Error
	}
	return cars, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
