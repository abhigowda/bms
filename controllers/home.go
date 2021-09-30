package controllers

import (
	"BMS/logic"
	model "BMS/models"
	"fmt"
	"net/http"
)

func GetDataForHome(r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside GetDataForHome")
	returnData.Code = 400

	movies, err := model.GetBmsMoviesByEmail()
	if err != nil {
		returnData.Msg = "Failed to read movies"
		return
	}
	returnMap := map[string]interface{}{
		"movies": movies,
	}
	fmt.Println("----------movies", movies)
	returnData.Msg = "Success"
	returnData.Code = 200
	returnData.Model = returnMap
	return
}
