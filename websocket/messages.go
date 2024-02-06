package wbs

import (
	"database/sql"
	"fmt"
	"net/http"
	"realtimeforum/models"

	"encoding/json"

	"github.com/gorilla/websocket"
)


func HandleCommunications(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	userFrom := r.URL.Query().Get("userFrom")
	toUser := r.URL.Query().Get("toUser")

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer ws.Close()

	err = models.MarkMessagesAsRead(db, toUser, userFrom)
	if err != nil {
		fmt.Println("error to udapte at message is read")
	}

	message := models.MessageSender{}
	user_forum, err := models.GetMessage(db, userFrom, toUser)
	user_receiver, err1 := models.GetMessage(db, toUser, userFrom)

	if err != nil || err1 != nil {
		fmt.Println("yo error")
		return
	}
	message.UserForum = user_forum
	message.UserReceiver = user_receiver
	jsonMsg, err := json.Marshal(message)
	if err != nil {
		fmt.Println("write:", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, jsonMsg)

	if err != nil {
		fmt.Println("write:", err)
		return
	}

}

