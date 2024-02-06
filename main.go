package main

import (
	"fmt"
	
	"net/http"
	"os"
	"realtimeforum/config"
	"realtimeforum/controllers"
	"realtimeforum/routes"	
	"time"
	 "github.com/rs/cors"
)

var (
	Port = ":8081"
)

const currentTime = "2006-01-02 15:04:05"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Purple  = "\033[95m"
	Dark    = "\033[90m"
)

// InitMessage prints a message when the server starts
func InitMessage() {
	fmt.Printf(Cyan + "===============================================\n" + Reset)
	fmt.Printf(Magenta + "Starting Realtime forum\n" + Reset)
	fmt.Printf(Magenta + "Server running on " + Green + "http://localhost" + Port + "\n" + Reset)
	fmt.Printf(Magenta + "Server started at: " + Blue + time.Now().Format(currentTime) + "\n" + Reset)
	fmt.Printf(Magenta + "Write" + Blue + " status" + Reset + Magenta + " to see loged in users\n" + Reset)
	fmt.Printf(Magenta + "Press Ctrl+C to stop the server\n" + Reset)
	fmt.Printf(Cyan + "===============================================\n" + Reset)
}

func init() {

	var err error

	controllers.DB, err = config.GetDB()
	if err != nil {
		fmt.Println("connection database Error")
		os.Exit(0)
	}

}




func main() {
	InitMessage()
	routes.Route()
	handler := cors.Default().Handler(http.DefaultServeMux)
	http.ListenAndServe(Port, handler)
	defer controllers.DB.Close()

}
