package model

// Room is the database model for a room
type Room struct {
	// DoorNumber is the primary key for the room. It is also the number of the room.
	DoorNumber int `gorm:"primaryKey"`
	// Description is the description of the room
	Description string `gorm:"not null"`
	// Floor is the floor of the room
	Floor int `gorm:"not null"`
	// Used for the many to many association with the user. The .Preload() method is used to load the reservations dynamically from GORM
	Reservations []Reservation
}
