package repository

import (
	"database/sql"
	"mini-project-quiz-3/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
	sql := "SELECT * FROM category ORDER BY id ASC"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(&category.ID, &category.Name, &category.Created_at, &category.Updated_at)
		if err != nil {
			panic(err)
		}
		result = append(result, category)
	}
	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	timeStamp := time.Now().Format(time.RFC3339)
	sql := "INSERT INTO category (name, created_at, updated_at) VALUES ($1, $2, $3)"
	errs := db.QueryRow(sql, category.Name, timeStamp, timeStamp)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	timeStamp := time.Now().Format(time.RFC3339)
	sql := "UPDATE category SET name = $1, updated_at = $2 WHERE id = $3"
	errs := db.QueryRow(sql, category.Name, timeStamp, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"
	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}

func GetBookByCategory(db *sql.DB, category structs.Category) (result []structs.Book, err error) {
	sql := "SELECT * FROM book WHERE category_id = $1 ORDER BY id ASC"
	rows, err := db.Query(sql, category.ID)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Created_at, &book.Updated_at, &book.Category_id)

		if err != nil {
			panic(err)
		}
		result = append(result, book)
	}

	return
}
