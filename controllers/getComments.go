package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"realtimeforum/models"
	"strconv"
	"strings"
)

func GetComments(w http.ResponseWriter, r *http.Request) {

	isAuth, _ := Auth(DB, w, r)
	if !isAuth{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
	}
	urlPath := r.URL.Path
	splitPath := strings.Split(urlPath, "/")
	id := splitPath[len(splitPath)-1]
	num, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("yes il y a get comments")
	}

	com := models.Comment{}
	comment, err := com.GetComments(DB, num)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
