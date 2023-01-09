package models

import "hopeugetknowuwont/validator"

func (u Signup) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("userid", u.UserID)
	v.Required("password", u.Password)

	return v.IsValid(), v.Errors
}

func (l Login) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("userid", l.UserID)
	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}
