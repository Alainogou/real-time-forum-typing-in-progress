package models

import (
	"database/sql"
	"fmt"
)

type Post struct {
	ID        int
	User_id   int
	Title     string
	Content   string
	ImageName string
	Category  []int
}
type UserData struct {
	Datas  interface{}
	IsAuth bool
	Cats   []Category
	// Pagin      models.Metadata
	User       User
	Comments   []Comment
	ErrorLog   string
	CurrentCat int
}

type AllPost struct {
	Post_id     int      `json:"Post_id"`
	Title       string   `json:"Title"`
	Content     string   `json:"Content"`
	ImageName   string   `json:"ImageName"`
	NickName    string   `json:"NickName"`
	User_id     int      `json:"User_id"`
	Nbrlike     int      `json:"Nbrlike"`
	NbrComments int      `json:"NbrComments"`
	Category    []string `json:"Category"`
}

func (post *Post) GetAllPosts(db *sql.DB) ([]AllPost, error) {
	Allpost := []AllPost{}
	var err error
	var row *sql.Rows

	req := `SELECT p.id, p.title, p.content, p.imgUrl, u.nickName, u.id,
					( SELECT count(*) FROM "user_post_reaction" "a" WHERE p.id=a."postId" AND isLiked) as "liked",
					( SELECT count(*) FROM "comment" "c" WHERE p.id=c."postId" ) as "Comments"
				FROM "Post" "p"
				JOIN "User" "u" ON p.userId = u.id ORDER BY p.id DESC;
				`
	row, err = db.Query(req)

	if err != nil {
		fmt.Println("eer", err)
		return []AllPost{}, err
	}

	for row.Next() {

		OnePosts := AllPost{}

		row.Scan(&OnePosts.Post_id, &OnePosts.Title, &OnePosts.Content, &OnePosts.ImageName, &OnePosts.NickName, &OnePosts.User_id, &OnePosts.Nbrlike, &OnePosts.NbrComments)
		category := Category{}
		err = category.GetCategory(db, OnePosts.Post_id)
		if err != nil {
			fmt.Println("err lors avec GetCAt")
		}
		OnePosts.Category = category.Name
		Allpost = append(Allpost, OnePosts)
	}
	return Allpost, row.Err()
}

func (post *Post) InsertPost(db *sql.DB) error {
	req := `INSERT INTO post ( userId,title,content, imgUrl) VALUES (?,?,?,?)`
	data, err := db.Exec(req, post.User_id, post.Title, post.Content, post.ImageName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	i, errres := data.LastInsertId()
	if errres != nil {
		fmt.Println(errres)
		return errres
	}
	for _, v := range post.Category {
		req2 := `INSERT INTO  category_relation ( postId, categoryId) VALUES (?,?);`
		_, errree := db.Exec(req2, i, v)
		if errree != nil {
			fmt.Println("errrese")
			return errree
		}
	}
	return nil

}

