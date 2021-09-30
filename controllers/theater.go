package controllers

import (
	"BMS/logic"
	model "BMS/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cast"
)

func GetTheaters(r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside GetTheaters")
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
	shows, err := model.GetTheaterByDateAndMovieId(inputObj["date"], cast.ToInt(inputObj["movieId"]))
	if err != nil {
		returnData.Msg = "Failed to read theaters"
		return
	}

	returnData.Msg = "Success"
	returnData.Code = 200
	returnData.Model = shows
	return
}

func CheckAvailableSeats(r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside CheckAvailableSeats")
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
	returnData.Model = map[string]string{
		"availableSeats": cast.ToString(theater.TotalCapacity - bookedSeats),
	}
	returnData.Msg = "Success"
	returnData.Code = 200
	return
}
