package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	wbs "realtimeforum/websocket"
	"strings"

	"encoding/json"
)

type UserDeconn struct {
	NickName string `json:"NickName"`
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	isAuth, _ := Auth(DB, w, r)
	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ok, _ := CheckRequest(r, "/logout", "post")

	if !ok {
		fmt.Println("errologout")
		return
	}

	deconn_user := UserDeconn{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(reqBody, &deconn_user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(deconn_user)

	session, err := r.Cookie("sessionid")

	errDelete := DeleteSession(DB, session.Value)
	if errDelete != nil || err != nil {
		fmt.Println("delete session error")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  session.Name,
		Value: "",
		Path:  session.Path,
	})
	wbs.RemoveUserFromMap(deconn_user.NickName, DB)
	fmt.Println("Decoonexion", wbs.UsersMap)

}

func DeleteSession(db *sql.DB, ssid string) error {
	req := `DELETE from sessions Where sessionId=?;`
	_, err := db.Exec(req, ssid)
	return err
}

func CheckRequest(r *http.Request, path, methode string) (bool, int) {
	if strings.ToLower(r.Method) == methode && r.URL.Path == path {
		return true, 0
	} else if !Getmethode(r, methode) {
		return false, 405
	} else {
		return false, 404
	}
}

func Getmethode(r *http.Request, methode string) bool {
	if strings.ToLower(r.Method) != methode {
		return false
	}
	return true
}
