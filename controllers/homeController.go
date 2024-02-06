
package controllers

import (

	"net/http"
	"log"
	"html/template"
	
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r) // Renvoie une page 404 si le chemin n'est pas "/"
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		// Loggez l'erreur et renvoyez une réponse d'erreur personnalisée
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		// Loggez l'erreur et renvoyez une réponse d'erreur personnalisée
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}