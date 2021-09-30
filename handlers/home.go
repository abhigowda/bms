package handlers

import (
	"BMS/controllers"
	"BMS/logic"
	"fmt"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside AuthHandler", r.URL.Path)
	switch true {
	case strings.Contains(r.URL.Path, "/home"):
		logic.URLReturnResponseJson(w, HandleHome(w, r))
	default:
		fmt.Println("In Default case")
	}
}

func HandleHome(w http.ResponseWriter, r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside HandleHome")
	// assigning the default response
	returnData.Msg = "Invalid session"
	returnData.Code = 400
	_, err := logic.IsAuthorized(w, r)
	if err != nil {
		return
	}
	switch r.Method {
	case http.MethodPost:
		returnData = controllers.GetDataForHome(r)
	default:
		fmt.Println("In default case :: method ::", r.Method)
	}
	return
}
