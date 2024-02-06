package helper

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

var u1 = uuid.Must(uuid.NewV4())

func SetCookieInDB(w http.ResponseWriter) string {
	sssid := u1.String() + "-" + time.Now().GoString()
	cookie := http.Cookie{
		Name:     "sessionid",
		Value:    sssid,
		Expires:  time.Now().Add(time.Minute * 30),
		Path:     "/",
		MaxAge:   3600 * 24 * 3,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
	return sssid
}
func SessionAddOrUpdate(db *sql.DB, sssid, useremail string) error {
	req := `SELECT sessionId,email, expires_at from sessions Where email='` + useremail + `';`
	// req:=fmt.Sprintf(`SELECT * from Session Where email=?;`)
	row, err := db.Query(req)
	var sessionid, email string
	var expires_at time.Time
	var errsession error
	if err != nil {
		fmt.Println(err)
		return err
	}

	for row.Next() {
		row.Scan(&sessionid, &email, &expires_at)

	}

	if email == useremail {
		_, errsession = db.Exec("UPDATE sessions SET sessionId=?,  expires_at=? where email=?;", sssid, time.Now().Add(time.Minute*30), email)
	} else {
		_, errsession = db.Exec("INSERT INTO sessions (sessionId,email, expires_at) VALUES(?,?,?);", sssid, useremail, time.Now().Add(time.Minute*30))
	}
	return errsession

}
