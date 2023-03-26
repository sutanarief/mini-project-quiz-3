package repository

import (
	"database/sql"
	"mini-project-quiz-3/helper"
	"mini-project-quiz-3/structs"
	"time"
)

func GetAllBook(db *sql.DB) (result []structs.Book, err error) {
	sql := "SELECT * FROM book ORDER BY id ASC"
	rows, err := db.Query(sql)

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

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	timeStamp := time.Now().Format(time.RFC3339)
	thickness := helper.ThicknessHandler(int(book.Total_page))
	isValidURL := helper.UrlValidator(book.Image_url)
	isValidYear := helper.YearValidator(book.Release_year)

	errInput := helper.ErrorHandler(isValidURL, isValidYear)

	if errInput != nil {
		return errInput
	}

	sql := "INSERT INTO book (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, thickness, timeStamp, timeStamp, book.Category_id)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	timeStamp := time.Now().Format(time.RFC3339)
	thickness := helper.ThicknessHandler(int(book.Total_page))
	isValidURL := helper.UrlValidator(book.Image_url)
	isValidYear := helper.YearValidator(book.Release_year)

	errInput := helper.ErrorHandler(isValidURL, isValidYear)

	if errInput != nil {
		return errInput
	}

	sql := "UPDATE book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, updated_at = $8, category_id = $9 WHERE id = $10"
	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, thickness, timeStamp, book.Category_id, book.ID)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id = $1"
	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}
