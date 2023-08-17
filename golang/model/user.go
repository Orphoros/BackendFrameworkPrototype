package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/mail"
)

// User is the database model for a user
type User struct {
	// UUID identifying the user
	Id string `gorm:"primaryKey"`
	// Name of the user
	Name string `gorm:"not null"`
	// Email of the user
	Email string `gorm:"not null"`
	// Password of the user
	Password string `gorm:"not null"`
	// Used for GORM to create the many to many association. The .Preload() method is used to load the reservations dynamically from GORM
	Reservations []Reservation
}

// UserLogin is a model for a user login request body
type UserLogin struct {
	Email    string
	Password string
}

// UserUpdate is a model for a user update request body
type UserUpdate struct {
	Name     string
	Email    string
	Password string
}

// FromHttpRequest is a method that parses the request body into the UserLogin struct. Returns an error if the body is not a valid json
func (u *UserUpdate) FromHttpRequest(body io.Reader) error {
	if body == nil {
		return fmt.Errorf("body is not defined")
	}
	jsonDecoder := json.NewDecoder(body)
	if jsonDecoder.Decode(u) != nil {
		return fmt.Errorf("body is not a valid json")
	}

	return nil
}

// FromHttpRequest is a method that parses the request body into the UserLogin struct. Returns an error if the body is not a valid json or if the email or password is not defined
func (u *User) FromHttpRequest(body io.Reader) error {
	if body == nil {
		return fmt.Errorf("body is not defined")
	}
	jsonDecoder := json.NewDecoder(body)
	if jsonDecoder.Decode(u) != nil {
		return fmt.Errorf("body is not a valid json")
	}

	if u.Name == "" {
		return fmt.Errorf("name is not defined")
	}

	if u.Password == "" {
		return fmt.Errorf("password is not defined")
	}

	if u.Email == "" {
		return fmt.Errorf("email is not defined")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("email is not valid")
	}

	if len(u.Name) >= 255 { // 255 is used to limit the MySQL varchar size
		return fmt.Errorf("username is too long")
	}

	if len(u.Password) >= 255 {
		return fmt.Errorf("password is too long")
	}

	return nil
}

// FromHttpRequest is a method that parses the request body into the UserLogin struct. Returns an error if the body is not a valid json or if the email or password is not defined
func (u *UserLogin) FromHttpRequest(body io.Reader) error {
	if body == nil {
		return fmt.Errorf("body is not defined")
	}
	jsonDecoder := json.NewDecoder(body)
	if jsonDecoder.Decode(u) != nil {
		return fmt.Errorf("body is not a valid json")
	}

	if u.Password == "" {
		return fmt.Errorf("password is not defined")
	}

	if u.Email == "" {
		return fmt.Errorf("email is not defined")
	}

	if _, err := mail.ParseAddress(u.Email); err != nil {
		return fmt.Errorf("email is not valid")
	}

	return nil
}
