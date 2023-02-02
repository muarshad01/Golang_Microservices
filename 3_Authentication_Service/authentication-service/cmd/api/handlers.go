package main

import (
	"errors"
	"fmt"
	"net/http"
)

// jsonPayload: From broker-service to authenticate-service
type jsonPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	
	//var requestPayload struct {
	//	Email    string `json:"email"`
	//	Password string `json:"password"`
	//}

	// read json into var
	var requestPayload jsonPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse {
		Error: false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data: user,
	}

	//var payload jsonResponse
	//payload.Error = true
	//payload.Message = fmt.Sprintf("Logged in user %s", user.Email)
	//payload.Data = user

	// jsonResponse: From authenticate-service to broker-service
	app.writeJSON(w, http.StatusAccepted, payload)
}