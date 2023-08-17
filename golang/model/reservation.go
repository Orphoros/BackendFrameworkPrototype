package model

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Reservation is the database model for a reservation. It functions as a join table between the User and Room tables
type Reservation struct {
	// Id is the primary key (UUID) for the reservation.
	Id string `gorm:"primaryKey"`
	// Date is the unix epoch date of the reservation
	Date int64
	// UserID is the foreign key for the User table. It is also the id of the user.
	UserID string `gorm:"size:40"`
	// RoomDoorNumber is the foreign key for the Room table. It is also the number of the room.
	RoomDoorNumber int
	// The User object is used to load the user data dynamically from GORM with Preload()
	User User `gorm:"foreignKey:UserID;references:Id"`
	// The Room object is used to load the room data dynamically from GORM with Preload()
	Room Room `gorm:"foreignKey:RoomDoorNumber;references:DoorNumber"`
}

// ReservationCreate is a model for a reservation create (POST) request body
type ReservationCreate struct {
	// Date is the date of the reservation. It must be in the format yyyy-mm-dd
	Date string
	// RoomNumber is the number of the room. It must be greater than 0 because 0 is the default, not given value
	RoomNumber int
}

// FromHttpRequest is a method that parses the request body into the ReservationCreate struct. Returns an error if the body is not a valid json, if the date is not in the format yyyy-mm-dd, if the date is in the past or today, or if the room number is not defined or is less than 1
func (r *ReservationCreate) FromHttpRequest(body io.Reader) error {
	if body == nil {
		return fmt.Errorf("body is not defined")
	}
	jsonDecoder := json.NewDecoder(body)
	if jsonDecoder.Decode(r) != nil {
		return fmt.Errorf("body is not a valid json")
	}

	t, err := time.Parse("2006-01-02", r.Date)
	if err != nil {
		return fmt.Errorf("date is not of the format yyyy-mm-dd")
	}

	if t.Before(time.Now()) {
		return fmt.Errorf("date cannot be in the past or today")
	}

	if r.RoomNumber < 0 {
		return fmt.Errorf("room number is not defined, or must be greater than 0")
	}

	return nil
}
