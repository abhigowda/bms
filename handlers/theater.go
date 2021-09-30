package handlers

import (
	"BMS/controllers"
	"BMS/logic"
	"fmt"
	"net/http"
	"strings"
)

func TheatersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside AuthHandler", r.URL.Path)
	switch true {
	case strings.Contains(r.URL.Path, "/theaters/checkSeats"):
		logic.URLReturnResponseJson(w, HandleCheckSeats(w, r))
	case strings.Contains(r.URL.Path, "/theaters"):
		logic.URLReturnResponseJson(w, HandleTheaters(w, r))
	default:
		fmt.Println("In Default case")
	}
}

func HandleTheaters(w http.ResponseWriter, r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside HandleTheaters")
	// assigning the default response
	returnData.Msg = "Invalid session"
	returnData.Code = 400
	_, err := logic.IsAuthorized(w, r)
	if err != nil {
		return
	}
	switch r.Method {
	case http.MethodGet:
		returnData = controllers.GetTheaters(r)
	default:
		fmt.Println("In default case :: method ::", r.Method)
	}
	return
}

func HandleCheckSeats(w http.ResponseWriter, r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside HandleCheckSeats")
	// assigning the default response
	returnData.Msg = "Invalid session"
	returnData.Code = 400
	_, err := logic.IsAuthorized(w, r)
	if err != nil {
		return
	}
	switch r.Method {
	case http.MethodGet:
		returnData = controllers.CheckAvailableSeats(r)
	default:
		fmt.Println("In default case :: method ::", r.Method)
	}
	return
}
