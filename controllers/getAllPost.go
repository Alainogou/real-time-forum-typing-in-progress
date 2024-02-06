package controllers

import (
	"encoding/json"
	"net/http"
	"realtimeforum/models"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	
	post := models.Post{}

	isAuth, _ := Auth(DB, w, r)
	if !isAuth{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
	}
	allpost, err := post.GetAllPosts(DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allpost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

