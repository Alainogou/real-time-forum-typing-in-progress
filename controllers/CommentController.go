package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"realtimeforum/models"
)

type CommentJson struct {
	UserId  int    `json:"UserId"`
	Post_id int    `json:"Post_id"`
	Content string `json:"Content"`
}

func CreateComment(w http.ResponseWriter, r *http.Request) {

	comments := CommentJson{}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(reqBody, &comments)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)

	}

	if len(comments.Content) == 0 {
		Error = "content is empty"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "emptycomment",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	com := models.Comment{}
	errinsert := com.InsertComments(DB, comments.Post_id, comments.UserId, comments.Content)

	if errinsert != nil {
		fmt.Println(errinsert)
		// helper.ErrorPage(w, 500)
		return
	}

}
