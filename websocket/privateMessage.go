package wbs

import (
	"database/sql"
	"fmt"
	"net/http"
	"realtimeforum/models"
	"time"

	"encoding/json"

	"github.com/gorilla/websocket"
)

type MessageJson struct {
	FromUser       string    `json:"FromUser"`
	ContentMessage string    `json:"Message"`
	ToUser         string    `json:"ToUser"`
	CreateDate     time.Time `json:"CreateDate"`
	ToUserClosed   string    `json:"ToUserClosed"`
	IsTyping	   bool  	 `json:"IsTyping"`
}

var MessageConnection = make(map[string]*websocket.Conn)
var PersonOpenChat = make(map[string]*websocket.Conn)

func HandlePrivateMessage(w http.ResponseWriter, r *http.Request, db *sql.DB) {

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

	for {
		MessageConnection[userFrom+toUser] = ws
		PersonOpenChat[userFrom+toUser] = ws

		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		var receivedMsg MessageJson
		
		err = json.Unmarshal(msg, &receivedMsg)
		if err != nil {
			fmt.Println("unmarshal:", err)
			return
		}
		// fmt.Println(receivedMsg, "avec typing")

		if receivedMsg.ToUserClosed != "" {

			delete(PersonOpenChat, receivedMsg.ToUserClosed)

			break
		}
		com := models.MessagePrivate{}

		if receivedMsg.ContentMessage != "" {

			errinsert := com.InsertMessage(db, receivedMsg.FromUser, receivedMsg.ToUser, receivedMsg.ContentMessage, receivedMsg.CreateDate)
			if errinsert != nil {
				fmt.Println(errinsert)
				return
			}
			Broad(userFrom, toUser, db)

		}else if receivedMsg.ToUserClosed ==""{
			
			if IsUserConnected(toUser + userFrom) {
				message := models.MessageSender{}
				user_forum, _:= models.GetMessage(db, toUser, userFrom)
				user_receiver, _ := models.GetMessage(db, userFrom, toUser)
				message.UserForum = user_forum
				message.UserReceiver = user_receiver
				message.IsTyping=receivedMsg.IsTyping

				jsonMsg, err := json.Marshal(message)
				if err != nil {
					fmt.Println("write:", err)
					return
				}
				conn := PersonOpenChat[toUser+userFrom]

				err = conn.WriteMessage(websocket.TextMessage, jsonMsg)

				if err != nil {
					fmt.Println("write:", err)
					return
				}

			}
		}

	}

}

func Broad(userFrom, toUser string, db *sql.DB) {

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
	conn := MessageConnection[userFrom+toUser]
	err = conn.WriteMessage(websocket.TextMessage, jsonMsg)

	if err != nil {
		fmt.Println("write:", err)
		return
	}

	if IsUserConnected(toUser + userFrom) {

		user_forum, _ = models.GetMessage(db, toUser, userFrom)
		user_receiver, _ = models.GetMessage(db, userFrom, toUser)
		message.UserForum = user_forum
		message.UserReceiver = user_receiver

		jsonMsg, err := json.Marshal(message)
		if err != nil {
			fmt.Println("write:", err)
			return
		}
		conn = PersonOpenChat[toUser+userFrom]

		err = conn.WriteMessage(websocket.TextMessage, jsonMsg)

		if err != nil {
			fmt.Println("write:", err)
			return
		}

	} else {
		if IsUserExist(toUser) {

			// for i := 0; i < len(UserSlice); i++ {

			// 	count, err := models.CountUnreadMessages(db, UserSlice[i].NickName, toUser)
			// 	if err != nil {
			// 		fmt.Println("error to give count read messages")

			// 	}
			// 	allUserStatus.AllUser[i].NbreMessages = count

			// }

			allUserS := AllUserStatus{}
			allUserS.NewMessage = true
			allUserS.PersonConnected = userFrom
			jsonAll, _ := json.Marshal(allUserS)
			conWith := UsersMap[toUser].Conn
			err = conWith.WriteMessage(websocket.TextMessage, jsonAll)
			if err != nil {
				fmt.Println("write:", err)
				return
			}
		}
	}

}

func IsUserConnected(username string) bool {
	_, exists := PersonOpenChat[username]
	return exists
}
