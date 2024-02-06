package models

import (
	"database/sql"
)

type Category struct {
	Name []string
}

func (Cat *Category) GetCategory(db *sql.DB, post_id int) error {
	req := `SELECT c.categoryName  from category c LEFT JOIN "category_relation" pt WHERE pt."categoryId"=c.id AND pt."postId"=?;`
	row, err := db.Query(req, post_id)

	if err != nil {
		return err
	}

	for row.Next() {
		var name string
		row.Scan(&name)
		Cat.Name = append(Cat.Name, name)

	}
	return row.Err()
}


