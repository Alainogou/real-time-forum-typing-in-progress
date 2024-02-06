package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"realtimeforum/models"
	"strings"
	"html"
	"time"
)

var Datas = models.UserData{}



type PostContent struct {
	User_id  int    `json:"User_id"`
	Title    string `json:"Title"`
	Content  string `json:"Content"`
	Category []int  `json:"Category"`
	Image    string `json:"Image"`
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	isAuth, _ := Auth(DB, w, r)
	if !isAuth{
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
	}

	post := models.Post{}

	var newPost PostContent
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(reqBody, &newPost)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if newPost.Title == "" {
		Error = " Please Enter a title"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "titleNoFound",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if len(newPost.Category) == 0 {
		Error = " Please choose a category"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "categoryNofound",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}
	if newPost.Content == "" {
		Error = " Please Enter your message"
		errorResponse := ErrorResponse{
			Message:    Error,
			ErrorClass: "contentNofound",
			Code:       http.StatusBadRequest,
		}
		sendReponseError(w, errorResponse, http.StatusBadRequest)
		return
	}

	if newPost.Image != "" {

		imageType := strings.SplitN(newPost.Image, ";", 2)[0][5:]
		ext := imageType[6:]

		if !CheckExtension(imageType) {
			Error = "image type incorrect (jpeg, jpg, png, gif)"
			errorResponse := ErrorResponse{
				Message:    Error,
				ErrorClass: "imageNoCorrect",
				Code:       http.StatusBadRequest,
			}
			sendReponseError(w, errorResponse, http.StatusBadRequest)
			return
		}

		base64Data := newPost.Image[strings.IndexByte(newPost.Image, ',')+1:]
		imageData, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			Error = "Error decoding image"
			errorResponse := ErrorResponse{
				Message:    Error,
				ErrorClass: "internalError",
				Code:       http.StatusBadRequest,
			}
			sendReponseError(w, errorResponse, http.StatusInternalServerError)
			return

		}

		if len(imageData)/1000000 >= 20 {
			Error = "image size >20 MO"
			errorResponse := ErrorResponse{
				Message:    Error,
				ErrorClass: "imageNoCorrect",
				Code:       http.StatusBadRequest,
			}
			sendReponseError(w, errorResponse, http.StatusBadRequest)
			return
		}

		// Supprimez les espaces blancs
		newPost.Image = strings.ReplaceAll(newPost.Image, " ", "")

		// Créez le chemin du fichier
		filename := fmt.Sprintf("%d-postImage.%s", time.Now().UnixNano(), ext)
		// Écrivez les données décodées dans un fichier
		err = ioutil.WriteFile("./assets/imageUpload/"+filename, imageData, 0644)
		if err != nil {
			Error = "Error saving image"
			errorResponse := ErrorResponse{
				Message:    Error,
				ErrorClass: "internalError",
				Code:       http.StatusBadRequest,
			}
			sendReponseError(w, errorResponse, http.StatusInternalServerError)
			return

		}

		post = models.Post{

			User_id:   newPost.User_id,
			Title:    html.EscapeString(strings.TrimSpace(newPost.Title)),
			Content:   html.EscapeString(strings.TrimSpace(newPost.Content)) ,
			ImageName: filename,
			Category:  newPost.Category,
		}

	} else {

		post = models.Post{
			User_id:  newPost.User_id,
			Title:    html.EscapeString(strings.TrimSpace(newPost.Title)),
			Content:   html.EscapeString(strings.TrimSpace(newPost.Content)) ,
			Category: newPost.Category,
		}

	}

	errpos := post.InsertPost(DB)

	if errpos != nil {
		fmt.Println(errpos)

		// helper.ErrorPage(w, 404)
		return

	}

}

func CheckExtension(ext string) bool {
	extensions := []string{"image/jpeg", "image/png", "image/gif", "image/jpg"}
	for _, validExt := range extensions {
		if strings.ToLower(ext) == validExt {
			return true
		}
	}
	return false
}
