package database

import (
	"UniDrive/back-end/api/models"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/gofrs/uuid"
)

func BookRide(db *gorm.DB, user_id string, ride_id string) (models.Booking, error) {
	var booking models.Booking
	rawUuid, err := uuid.NewV4()
	if err != nil {
		return booking, err
	}
	uuid := rawUuid.String()
	
	
	timestamp := time.Now().Format(time.RFC3339)

	result := db.Exec("INSERT INTO booking(id, ride_id, booking_timestamp, passenger_id) values (?,?,?,?) ",uuid,ride_id,timestamp,user_id)
	if result.Error != nil {
		return booking, result.Error
	}
	

	var carDetails models.CarDetails

	result = db.Raw("SELECT * FROM car_details WHERE user_id = (SELECT driver_id FROM ride WHERE id = ?)",ride_id).Find(&carDetails)
	if result.Error != nil {
		return booking, result.Error
	}

	result = db.Exec("UPDATE ride SET available_seats = available_seats - 1   WHERE id = ?", ride_id)
	if result.Error != nil {
		return booking, result.Error
	}
	
	booking.RideId = ride_id
	booking.BookingTimestamp = timestamp
	booking.CarDetails = carDetails
	
	return booking, nil
}
