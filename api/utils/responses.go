package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func Message(status bool, message string) map[string]interface{} {
return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, statusCode int, data interface{})  {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_ = json.NewEncoder(w).Encode(data)
}

func ERROR (w http.ResponseWriter, statusCode int, err error){
	w.WriteHeader(statusCode)
	if err != nil {
		fmt.Print(err.Error())
	}
}

func ErrorType(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("email Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect Password")
	}
	return errors.New("incorrect Details")
}