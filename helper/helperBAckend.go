package helper

import (
	"database/sql"
	"fmt"
	"realtimeforum/models"

	// "forum/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Data struct {
	Datas     interface{}
	IsAuth    bool
	
	User      models.User
	CurrenCat int
	ErrorLog  string
}

// func SessionAddOrUpdate(db *sql.DB, sssid, useremail string) error {
// 	req := `SELECT sessionId,email,datefin from Session Where email='` + useremail + `';`
// 	// req:=fmt.Sprintf(`SELECT * from Session Where email=?;`)
// 	row, err := db.Query(req)
// 	var sessionid, email string
// 	var datef time.Time
// 	var errsession error
// 	if err != nil {func CheckRequest(r *http.Request, path, methode string) (bool, int) {
// 		if strings.ToLower(r.Method) == methode && r.URL.Path == path {
// 			return true, 0
// 		} else if !Getmethode(r, methode) {
// 			return false, 405
// 		} else {
// 			return false, 404
// 		}
// 	}
// 		fmt.Println(err)
// 		return err
// 	}

// 	for row.Next() {
// 		row.Scan(&sessionid, &email, &datef)

// 	}

// 	if email == useremail {
// 		_, errsession = db.Exec("UPDATE Session SET sessionId=?, datefin=? where email=?;", sssid, time.Now().Add(time.Hour*24*3), email)
// 	} else {
// 		_, errsession = db.Exec("INSERT INTO Session (sessionId,email,datefin) VALUES(?,?,?);", sssid, useremail, time.Now().Add(time.Hour*24*3))
// 	}
// 	return errsession

// }

func Auth(Db *sql.DB, r *http.Request) (bool, string) {

	sessionpi, err := r.Cookie("sessionid")
	if err != nil || sessionpi.String() == "" {
		return false, ""
	}
	var Id int
	var sessionId, email string
	var datef time.Time
	req := `SELECT * from Session Where sessionId=?;`
	row, err := Db.Query(req, sessionpi.Value)

	if err != nil {
		return false, ""
	}
	for row.Next() {
		row.Scan(&Id, &sessionId, &email, &datef)
	}

	if sessionId != "" && email != "" && datef.After(time.Now()) {
		return true, email
	}
	return false, ""
}
func CheckRequest(r *http.Request, path, methode string) (bool, int) {
	if strings.ToLower(r.Method) == methode && r.URL.Path == path {
		return true, 0
	} else if !Getmethode(r, methode) {
		return false, 405
	} else {
		return false, 404
	}
}
func Getmethode(r *http.Request, methode string) bool {
	if strings.ToLower(r.Method) != methode {
		return false
	}
	return true
}
func DeleteSessio(db *sql.DB, ssid string) error {
	req := `DELETE from Session Where sessionId=?;`
	_, err := db.Exec(req, ssid)
	return err
}

// ******************************* PARSE FILE IN URL *****************
func PArseUlr(r *http.Request, match string) (bool, int) {
	index := strings.Split(r.URL.Path[1:], "/")
	if len(index) == 2 && index[0] == match {
		id, err := strconv.Atoi(index[1])
		if err == nil {
			return true, id
		}
	}
	return false, 0
}

func FecthError(ch []error) bool {
	for _, err := range ch {
		if err != nil {
			fmt.Println(err)
			return true
		}
	}
	return false
}

func ParseCatId(cat []string) ([]int, error) {
	catid := []int{}
	for _, v := range cat {
		a, errt := strconv.Atoi(v)
		if errt != nil {
			return []int{}, errt
		}
		catid = append(catid, a)
	}
	return catid, nil
}

// func GetAllPosts(db *sql.DB) ([]models.AllPost, error) {
// 	allpost := []models.AllPost{}
// 	var err error
// 	var row *sql.Rows

// 	req := `SELECT p.id, p.title, p.content, p.imgUrl, u.nickName,
// 					( SELECT count(*) FROM "user_post_reaction" "a" WHERE p.id=a."postId" AND isLiked) as "liked",
// 					( SELECT count(*) FROM "comment" "c" WHERE p.id=c."postId" ) as "Comments"
// 				FROM "Post" "p"
// 				JOIN "User" "u" ON p.userId = u.id ORDER BY p.id DESC;
// 				`
// 	row, err = db.Query(req)

// 	if err != nil {
// 		fmt.Println("eer", err)
// 		return allpost, err
// 	}

// 	for row.Next() {

// 		user := models.User{}
// 		OnePosts := models.AllPost{Poster: user, OnePost: models.Post{}}

// 		row.Scan(&OnePosts.OnePost.ID, &OnePosts.OnePost.Title, &OnePosts.OnePost.Content, &OnePosts.OnePost.ImageName, &OnePosts.Poster.NickName, &OnePosts.Nbrlike, &OnePosts.NbrComments)

// 		category := models.Category{}
// 		err = category.GetCategory(db, OnePosts.OnePost.ID)
// 		if err != nil {
// 			fmt.Println("err lors avec GetCAt")
// 		}
// 		OnePosts.Category = category.Name
// 		allpost = append(allpost, OnePosts)
// 	}
// 	return allpost, row.Err()
// }

// func GetData(r *http.Request, db *sql.DB, f func(*sql.DB, models.Pagination, string) ([]models.AllPost, error), pagination models.Pagination, w http.ResponseWriter, isAuth bool, metadata models.Metadata, user models.User) (Data, error) {
// 	var category models.Category
// 	// CatPost:=models.CatPost{}
// 	Cat, errcookie := r.Cookie("cat")
// 	Cats := ""
// 	if errcookie == nil {
// 		Cats = Cat.Value
// 	}
// 	data, errs := f(db, pagination, Cats)
// 	if errs != nil {
// 		return Data{}, errs
// 	}

// 	categories, errc := category.GetCategory(db)
// 	if errc != nil {
// 		return Data{}, errc
// 	}
// 	Data := Data{Datas: data, IsAuth: isAuth, Cats: categories, Pagin: metadata, User: user}
// 	return Data, nil
// }

// func SetPagination(db *sql.DB, r *http.Request, user models.User, query string) (models.Pagination, models.Metadata, error) {
// 	pageParam := r.URL.Query().Get("page")
// 	if pageParam == "" {
// 		pageParam = "1"
// 	}
// 	var err error
// 	models.ActualPage, err = strconv.Atoi(pageParam)
// 	if err != nil || models.ActualPage <= 0 {
// 		models.ActualPage = 1
// 	}
// 	pagination := models.Pagination{
// 		PageSize: 6,
// 		Page:     models.ActualPage,
// 	}
// 	totalRecords, err := models.GetTotalRecords(query, user, db)
// 	if err != nil {
// 		return models.Pagination{}, models.Metadata{}, err
// 	}
// 	metadata := models.GetMetadata(totalRecords, pagination.Page, pagination.PageSize)
// 	if pagination.Page > metadata.LastPage {
// 		pagination.Page = metadata.LastPage
// 		metadata.CurrentPage = pagination.Page
// 	}
// 	return pagination, metadata, nil
// }

// func GetFilterCat(ListPost_id []int, posts []models.AllPost) []models.AllPost {
// 	FilterPosts := []models.AllPost{}
// 	if len(ListPost_id) > 0 {
// 		for _, v := range posts {
// 			fmt.Println(v.OnePost.ID)
// 			for _, y := range ListPost_id {
// 				if v.OnePost.ID == y {
// 					FilterPosts = append(FilterPosts, v)
// 					break
// 				}
// 			}
// 		}
// 		return FilterPosts
// 	}
// 	return posts
// }

// func List_posts_id(db *sql.DB, cat_id string) []int {
// 	caId, err := strconv.Atoi(cat_id)
// 	if err != nil {
// 		return []int{}
// 	}
// 	categorie := models.Category{}
// 	ListPost_id, errPost := categorie.Post_id(db, caId)
// 	if errPost != nil {
// 		return []int{}
// 	}
// 	return ListPost_id
// }
