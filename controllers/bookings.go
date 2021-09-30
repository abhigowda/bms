package controllers

import (
	"BMS/logic"
	model "BMS/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

func AddBooking(r *http.Request, userId int) (returnData logic.ResponseJSON) {
	fmt.Println("Inside AddBooking")
	returnData.Code = 400
	inputObj := make(map[string]string)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnData.Msg = "Invalid Request"
		return
	}
	err = json.Unmarshal(body, &inputObj)
	if err != nil {
		returnData.Msg = "Invalid Request"
		return
	}
	// get theater by ID
	theater, err := model.GetBmsTheaterById(cast.ToInt(inputObj["theaterId"]))
	if err != nil {
		returnData.Msg = "Error while reading theater data"
		return
	}
	// get total seats booked on date, time ,movie and theater
	bookedSeats, err := model.GetTotalBooked(inputObj["showDate"], inputObj["showTime"], cast.ToInt(inputObj["theaterId"]), cast.ToInt(inputObj["movieId"]))
	if err != nil {
		returnData.Msg = "Error while reading booked seats"
		return
	}

	if theater.TotalCapacity <= bookedSeats {
		returnData.Msg = "All seats are booked"
		return
	}

	if bookedSeats+cast.ToInt(inputObj["seats"]) > theater.TotalCapacity {
		returnData.Msg = "Only " + cast.ToString(theater.TotalCapacity-bookedSeats) + " seats are available"
		return
	}
	showDate, _ := time.Parse("2006-01-02", inputObj["showDate"])
	book := model.BmsBookings{
		UserId:     userId,
		MovieId:    cast.ToInt(inputObj["movieId"]),
		TheaterId:  cast.ToInt(inputObj["theaterId"]),
		ShowDate:   showDate,
		ShowTime:   inputObj["showTime"],
		TotalSeats: cast.ToInt(inputObj["seats"]),
	}

	id, err := model.AddBmsBookings(&book)
	if err != nil {
		returnData.Msg = "Error while Booking"
		return
	}
	returnData.Model = map[string]string{
		"bookingId": cast.ToString(id),
	}
	returnData.Msg = "Success"
	returnData.Code = 200
	return
}

func GetBooking(r *http.Request, userId int) (returnData logic.ResponseJSON) {
	fmt.Println("Inside AddBooking")
	returnData.Code = 400

	bookings, err := model.GetBookingBYuserId(userId)
	if err != nil {
		returnData.Msg = "No Booking available"
		return
	}
	returnData.Model = bookings

	returnData.Msg = "Success"
	returnData.Code = 200
	return
}
