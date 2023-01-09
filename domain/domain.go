package domain

import (
	"errors"
	"hopeugetknowuwont/postgres"
)

var (
	ErrBadCredentials  = errors.New("userid/password combination is wrong")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UsersRepo   postgres.UsersRepo
	GarageRepos postgres.GarageRepo
}

func NewDomain(userRepo postgres.UsersRepo, garRepo postgres.GarageRepo) *Domain {
	return &Domain{UsersRepo: userRepo, GarageRepos: garRepo}
}
