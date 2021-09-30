package logic

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(id int, email string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func IsAuthorized(w http.ResponseWriter, r *http.Request) (id int, err error) {
	if r.Header["Token"] == nil {
		err = errors.New("token is empty")
		return
	}

	var mySigningKey = []byte(Secretkey)

	token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		err = errors.New("token has expired")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id = cast.ToInt(claims["id"])
		return
	}
	err = errors.New("error while extrating ID")
	return
}
