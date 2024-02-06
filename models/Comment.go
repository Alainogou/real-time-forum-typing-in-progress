package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID       int
	Content  string
	NickName string
}

func (Com *Comment) GetComments(db *sql.DB, post_id int) ([]Comment, error) {

	req := `SELECT c.id,c.content,u.Nickname
	FROM "Comment" c 
	INNER JOIN "User" u on c."userId"=u.id  WHERE c."postId"=?
	ORDER BY c."created_at" DESC;`

	comments := []Comment{}

	row, err := db.Query(req, post_id)
	if err != nil {
		return comments, err
	}
	for row.Next() {

		comment := Comment{}
		row.Scan(&comment.ID, &comment.Content, &comment.NickName)
		comments = append(comments, comment)
	}
	return comments, row.Err()
}

func (Com *Comment) InsertComments(db *sql.DB, post_id, user_id int, content string) error {
	req := `INSERT INTO comment (userId,postId,content,created_at) VALUES(?,?,?,?);`
	_, errr := db.Exec(req, user_id, post_id, content, time.Now())
	return errr
}
