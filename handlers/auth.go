package handlers

import (
	"BMS/controllers"
	"BMS/logic"
	"fmt"
	"net/http"
	"strings"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside AuthHandler", r.URL.Path)
	switch true {
	case strings.Contains(r.URL.Path, "/signup"):
		logic.URLReturnResponseJson(w, HandleSignUp(w, r))
	case strings.Contains(r.URL.Path, "/login"):
		logic.URLReturnResponseJson(w, HandleLogin(w, r))
	default:
		fmt.Println("In Default case")
	}
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside HandleSignUp")
	// assigning the default response
	returnData.Msg = "Invalid session"
	returnData.Code = 400

	switch r.Method {
	case http.MethodPost:
		returnData = controllers.SignUp(r)
	default:
		fmt.Println("In default case :: method ::", r.Method)
	}
	return
}

func HandleLogin(w http.ResponseWriter, r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside HandleLogin")
	// assigning the default response
	returnData.Msg = "Invalid session"
	returnData.Code = 400

	switch r.Method {
	case http.MethodGet:
		returnData = controllers.Login(r)
	default:
		fmt.Println("In default case :: method ::", r.Method)
	}
	return
}
