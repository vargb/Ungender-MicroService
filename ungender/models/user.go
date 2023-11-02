package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   string `json:"userid" gorm:"primaryKey"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Phno     string `json:"phno"`
	Password string `json:"password"`
	CarID    string `gorm:"foreignKey:CarID"`
}

func (g *Garage) InUseNow(user *User) bool {
	if g.CarID == user.UserID {
		return g.Available
	}
	return g.Available
}

func (u *User) HashPassword(password string) error {
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

// func (u *User) GenToken() (*AuthToken, error) {
// 	//15 Min Expiration
// 	expiredAt := time.Now().Add(time.Minute * 15)

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		ExpiresAt: expiredAt.Unix(),
// 		Id:        u.UserID,
// 		IssuedAt:  time.Now().Unix(),
// 		Issuer:    "dubai-flyin",
// 	})

// 	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &AuthToken{
// 		AccessToken: accessToken,
// 		ExpiredAt:   expiredAt,
// 	}, nil
// }

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
