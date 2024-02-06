package controllers

import (
	// "fmt"
	"database/sql"
	"html"
	"io/ioutil"
	"net/http"
	"regexp"

	"strconv"

	"encoding/json"
	"realtimeforum/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var Error string

var DB *sql.DB

// var user = models.User{}
type ErrorResponse struct {
	ErrorClass string `json:"error_class"`
	Message    string `json:"message"`
	Code       int    `json:"code"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	newUser := models.User{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(reqBody, &newUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if len(newUser.NickName) < 3 || len(newUser.NickName) > 10 || isEmailValid(strings.ToLower(newUser.NickName)) {
		Error = "Please enter a valid nickname (3 to 5 characters)"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "errNickname",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if newUser.Age == 0 {
		Error = "Invalid age"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "errAge",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if strings.ToLower(newUser.Gender) != "male" && strings.ToLower(newUser.Gender) != "female" {
		Error = "Invalid gender"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "errGender",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return

	}
	if !verifLen(newUser.LastName) {
		Error = "Invalid lastname"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "errLastName",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if !verifLen(newUser.FirstName) {
		Error = "Invalid FirstName"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "errFirstName",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if !isEmailValid(newUser.Email) {
		Error = "bad format of email"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "errEmail",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if newUser.Password != newUser.ConfirmPassword {
		Error = "Passwords do not match"
		errorResponse := ErrorResponse{
			// Message: Error,
			ErrorClass: "errPassword",
			Message:    Error,
			Code:       http.StatusBadRequest,
		}

		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	} else {
		haspassword, erft := bcrypt.GenerateFromPassword([]byte(newUser.Password), 5)
		newUser.Password = string(haspassword)
		newUser.FirstName = html.EscapeString(strings.TrimSpace(newUser.FirstName))
		newUser.LastName = html.EscapeString(strings.TrimSpace(newUser.LastName))
		newUser.Email = html.EscapeString(strings.TrimSpace(newUser.Email))
		newUser.NickName = html.EscapeString(strings.TrimSpace(newUser.NickName))
		newUser.Gender = html.EscapeString(strings.TrimSpace(newUser.Gender))
		err = newUser.InsertData(DB, newUser.Age, newUser.NickName, newUser.Email, newUser.LastName, newUser.FirstName, newUser.Password, newUser.Gender)

		if err != nil || erft != nil {
			if strings.HasPrefix(err.Error(), "UNIQUE constraint failed:") {
				Error = "Email or NickName already  exists"
				errorResponse := ErrorResponse{
					ErrorClass: "errEmailorNickname",
					Message:    Error,
					Code:       http.StatusBadRequest,
				}
				sendReponseError(w, errorResponse, http.StatusBadRequest)
				return
			} else {
				errorResponse := ErrorResponse{
					Message: Error,
					Code:    http.StatusInternalServerError,
				}
				sendReponseError(w, errorResponse, http.StatusInternalServerError)
				return
			}
		}
		models.IsNewUser = true

	}

}

func sendReponseError(w http.ResponseWriter, errorResponse interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)

}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func verifLen(data string) bool {
	return data != "" && !isNumber(data)
}
