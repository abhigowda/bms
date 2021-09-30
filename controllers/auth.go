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

type signUpData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const SecretKey = "secret"

func SignUp(r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside SignUp")

	returnData.Code = 400

	inputObj := signUpData{}
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
	// encrypt password
	password, err := logic.GeneratehashPassword(inputObj.Password)
	if err != nil {
		returnData.Msg = "Error while hashing password"
		return
	}

	user := model.BmsUser{
		Name:     inputObj.Name,
		Email:    inputObj.Email,
		Password: cast.ToString(password),
		Gender:   inputObj.Gender,
		Mobile:   inputObj.Mobile,
	}

	_, err = model.AddBmsUser(&user)
	if err != nil {
		returnData.Msg = "DB Error"
		return
	}

	returnData.Msg = "Success"
	returnData.Code = 200
	return
}

func Login(r *http.Request) (returnData logic.ResponseJSON) {
	fmt.Println("Inside Login")

	returnData.Code = 400

	inputObj := LoginData{}
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
	// read user by email
	user, err := model.GetBmsUserByEmail(inputObj.Email)
	if err != nil {
		returnData.Msg = "Error while fetching user data"
		return
	}

	if ok := logic.CheckPasswordHash(inputObj.Password, user.Password); !ok {
		returnData.Msg = "Invalid Password"
		return
	}
	// generate jwt token
	token, err := logic.GenerateJWT(user.Id, user.Email)
	if err != nil {
		returnData.Msg = "Failed to generate token"
		return
	}

	returnMap := map[string]string{
		"token": token,
	}

	returnData.Msg = "Success"
	returnData.Code = 200
	returnData.Model = returnMap
	return
}
