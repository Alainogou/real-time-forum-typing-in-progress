package models

import (
	"database/sql"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

var IsNewUser bool

type UserWithLastMessage struct {
	*User
	LastMessageAt time.Time
}

type User struct {
	Id              int             `json:"Id"`
	NickName        string          `json:"NickName"`
	Email           string          `json:"Email"`
	LastName        string          `json:"LastName"`
	FirstName       string          `json:"FirstName"`
	Password        string          `json:"Password"`
	Age             int             `json:"Age"`
	Gender          string          `json:"Gender"`
	ConfirmPassword string          `json:"ConfirmPassword"`
	Status          string          `json:"Status"`
	Conn            *websocket.Conn `json:"Conn"`
	LastMessageAt   sql.NullTime    `json:"LastMessageAt"` // Ajouté pour stocker la date du dernier message

}

type Session struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ExpiresAt time.Time
	CreatedAt time.Time
}

var Table = "user"

func (us *User) InsertData(db *sql.DB, age int, data ...string) error {
	req := fmt.Sprintf("INSERT INTO %s (age, nickName, email, lastName, firstName, password,  gender) VALUES(%d, '%s')", Table, age, strings.Join(data, "','"))
	dat, err := db.Prepare(req)
	if err != nil {
		return err
	}
	result, er := dat.Exec()
	fmt.Println(result)
	return er
}

func (UserOne *User) GetOneUser(db *sql.DB, email string) error {
	req := `SELECT id, age, nickName, email, lastName, firstName, password,  gender from ` + Table + ` Where email=?;`
	row := db.QueryRow(req, email)

	err := row.Scan(&UserOne.Id, &UserOne.Age, &UserOne.NickName, &UserOne.Email, &UserOne.LastName, &UserOne.FirstName, &UserOne.Password, &UserOne.Gender)

	UserOne.Email = html.UnescapeString(UserOne.Email)
	UserOne.NickName = html.UnescapeString(UserOne.NickName)
	UserOne.Password = html.UnescapeString(UserOne.Password)
	UserOne.LastName = html.UnescapeString(UserOne.LastName)
	UserOne.FirstName = html.UnescapeString(UserOne.FirstName)

	return err
}

func (UserOne *User) GetOneUserWithNickName(db *sql.DB, nickName string) error {
	req := `SELECT id, age, nickName, email, lastName, firstName, password,  gender from ` + Table + ` Where nickName=?;`
	row := db.QueryRow(req, nickName)

	err := row.Scan(&UserOne.Id, &UserOne.Age, &UserOne.NickName, &UserOne.Email, &UserOne.LastName, &UserOne.FirstName, &UserOne.Password, &UserOne.Gender)

	UserOne.Email = html.UnescapeString(UserOne.Email)
	UserOne.NickName = html.UnescapeString(UserOne.NickName)
	UserOne.Password = html.UnescapeString(UserOne.Password)
	UserOne.LastName = html.UnescapeString(UserOne.LastName)
	UserOne.FirstName = html.UnescapeString(UserOne.FirstName)

	return err
}

func GetAllUser(db *sql.DB) ([]*User, error) {
	req := `SELECT id, age, nickName, email, lastName, firstName, password, gender FROM ` + Table
	rows, err := db.Query(req)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Age, &user.NickName, &user.Email, &user.LastName, &user.FirstName, &user.Password, &user.Gender)
		if err != nil {
			return nil, err
		}

		user.Email = html.UnescapeString(user.Email)
		user.NickName = html.UnescapeString(user.NickName)
		user.Password = html.UnescapeString(user.Password)
		user.LastName = html.UnescapeString(user.LastName)
		user.FirstName = html.UnescapeString(user.FirstName)

		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func IsUserExist(db *sql.DB, email string) (string, error) {
	var userEmail string

	query := "SELECT email FROM User WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&userEmail)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", err
		}
	}

	return userEmail, nil
}
