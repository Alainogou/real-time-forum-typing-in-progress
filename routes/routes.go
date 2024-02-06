package routes

import (

	"net/http"
	
	"realtimeforum/controllers"
	wbs "realtimeforum/websocket"
	
	// "github.com/rs/cors"
)

func Route() {

	http.Handle("/assets/", http.StripPrefix("/assets/",  http.FileServer(http.Dir("./assets/"))))
	http.Handle("/css/", http.StripPrefix("/css/",  http.FileServer(http.Dir("./css/"))))
	http.Handle("/js/", http.StripPrefix("/js/",  http.FileServer(http.Dir("./js/"))))
	http.HandleFunc("/", controllers.HomeHandler)
	http.HandleFunc("/register", controllers.RegisterUser)
	http.HandleFunc("/login", controllers.LoginUser)
	http.HandleFunc("/auth", controllers.IsAuth)
	http.HandleFunc("/logout", controllers.LogoutUser)
	http.HandleFunc("/createPost", controllers.CreatePost)
	http.HandleFunc("/fetchPost", controllers.GetPosts)
	http.HandleFunc("/createComment", controllers.CreateComment)
	http.HandleFunc("/fetchComment/", controllers.GetComments)
	
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wbs.HandleConnections(w, r, controllers.DB)
	})
	http.HandleFunc("/communication", func(w http.ResponseWriter, r *http.Request) {
		wbs.HandleCommunications(w, r, controllers.DB)
	})

	http.HandleFunc("/privateSocket", func(w http.ResponseWriter, r *http.Request) {
		wbs.HandlePrivateMessage(w, r, controllers.DB)
	})

	
}