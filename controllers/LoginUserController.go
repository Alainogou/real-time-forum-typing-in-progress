package controllers

import (
	// "fmt"

	"fmt"
	"html"
	"net/http"
	"realtimeforum/helper"
	"realtimeforum/models"

	"io/ioutil"

	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	EmailOrUsername string `json:"EmailOrUsername"`
	Password        string `json:"Password"`
}

// var u1 = uuid.Must(uuid.NewV4())
// var Datas = UserData{}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	newLog := LoginRequest{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(reqBody, &newLog)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user := models.User{}

	if isEmailValid(newLog.EmailOrUsername) {
		err = user.GetOneUser(DB, html.EscapeString(newLog.EmailOrUsername))
	} else {
		err = user.GetOneUserWithNickName(DB, html.EscapeString(newLog.EmailOrUsername))

	}

	if err != nil {
		fmt.Println(err)

		Error = "You don't have account"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "logNotMatch",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return

	}

	ispassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newLog.Password))

	if ispassword != nil {
		Error = "Password is incorrect"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "logNotMatch",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	} else {

		sssid := helper.SetCookieInDB(w)
		errss := helper.SessionAddOrUpdate(DB, sssid, user.Email)
		if errss != nil {
			fmt.Println(errss)
			// helper.ErrorPage(w, 500)
			return
		}
		// Datas.ErrorLog = ""
		// http.Redirect(w, r, "/", 302)
		sssid = ""

		return
	}

}
