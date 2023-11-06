package potgres

import (
	"hopeugetknowuwont/ungender/graph/model"

	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	Login model.Login
}

type Claims struct {
	Phno string `json:"phno"`
	jwt.RegisteredClaims
}

type User struct {
	Phno     string `json:"phno" gorm:"primaryKey;unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	UserId   string `json:"userid"`
	CarId    string `json:"carid" gorm:"foreignKey:Carid"`
}

type Garage struct {
	Carid           string `json:"Carid" gorm:"primaryKey;unique;not null"`
	DateOfManufac   string `json:"DateOfManufac"`
	LastServiceDate string `json:"LastServiceDate"`
	LastUsedDate    string `json:"LastUsedDate" gorm:"not null"`
	Available       bool   `json:"Available" gorm:"not null"`
}
