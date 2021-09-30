package handlers

import (
	"BMS/controllers"
	"BMS/logic"
	"fmt"
	"net/http"
	"strings"
)

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside AuthHandler", r.URL.Path)
	switch true {
	case strings.Contains(r.URL.Path, "/booking"):
		logic.URLReturnResponseJson(w, HandleBooking(w, r))
	default:
		fmt.Println("In Default case")
	}
}

func HandleBooking(w http.ResponseWriter, r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside HandleBooking")
	// assigning the default response
	returnData.Msg = "Invalid session"
	returnData.Code = 400
	userId, err := logic.IsAuthorized(w, r)
	if err != nil {
		return
	}
	switch r.Method {
	case http.MethodPost:
		returnData = controllers.AddBooking(r, userId)
	case http.MethodGet:
		returnData = controllers.GetBooking(r, userId)
	default:
		fmt.Println("In default case :: method ::", r.Method)
	}
	return
}
