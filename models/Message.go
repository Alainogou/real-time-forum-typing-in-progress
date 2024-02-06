package models

import (
	"database/sql"
	"errors"
	"time"
)

type MessagePrivate struct {
	ID             int
	FromUser       string
	ContentMessage string
	ToUser         string
	CreateDate     time.Time
}

type MessageSender struct {
	UserForum    []MessagePrivate
	UserReceiver []MessagePrivate
	NewMessage   bool
	IsTyping 	 bool
}

func GetMessage(db *sql.DB, userFrom, toUser string) ([]MessagePrivate, error) {

	req := `SELECT m.id,m.message,m.createdDate
	FROM message m
	WHERE m."fromUser"=? and m."toUser"=?
	ORDER BY m."createdDate" ASC;`

	allmessage := []MessagePrivate{}

	row, err := db.Query(req, userFrom, toUser)
	if err != nil {
		return allmessage, err
	}
	for row.Next() {

		message := MessagePrivate{}
		row.Scan(&message.ID, &message.ContentMessage, &message.CreateDate)
		message.ToUser = toUser
		message.FromUser = userFrom
		allmessage = append(allmessage, message)

	}
	return allmessage, row.Err()
}

func (Com *MessagePrivate) InsertMessage(db *sql.DB, fromUser, toUser, contentMessage string, createDate time.Time) error {
	req := `INSERT INTO message (fromUser,toUser,message,isRead,createdDate) VALUES(?,?,?,?,?);`
	_, errr := db.Exec(req, fromUser, toUser, contentMessage, 0, time.Now())
	return errr
}

func NombreMessage(db *sql.DB, userFrom, toUser string) (int, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM message WHERE (fromUser = ? AND toUser = ?) OR (toUser = ? AND fromUser = ?)`, userFrom, toUser, userFrom, toUser).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func MarkMessagesAsRead(db *sql.DB, userFrom, toUser string) error {
	_, err := db.Exec(`UPDATE message SET isRead = 1 WHERE fromUser = ? AND toUser = ? AND isRead = 0`, userFrom, toUser)
	return err
}
func CountUnreadMessages(db *sql.DB, userFrom, toUser string) (int, error) {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM message WHERE fromUser = ? AND toUser = ? AND isRead = 0`, userFrom, toUser).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetLastMessageDateBetweenUsers récupère la date de création du dernier message entre deux utilisateurs.
func GetLastMessageDateBetweenUsers(db *sql.DB, userFrom string, toUser string) (*time.Time, error) {
	var lastMessageDate time.Time
	query := `
		SELECT createdDate
		FROM message
		WHERE (fromUser = ? AND toUser = ?) OR (fromUser = ? AND toUser = ?)
		ORDER BY createdDate DESC
		LIMIT 1
	`
	err := db.QueryRow(query, userFrom, toUser, toUser, userFrom).Scan(&lastMessageDate)
	if err != nil {
		if err == sql.ErrNoRows {
			// Pas de messages trouvés entre ces deux utilisateurs
			return nil, errors.New("no messages found between users")
		}
		return nil, err
	}
	return &lastMessageDate, nil
}
