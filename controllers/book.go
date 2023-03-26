package controllers

import (
	"mini-project-quiz-3/database"
	"mini-project-quiz-3/repository"
	"mini-project-quiz-3/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"result": "Success Insert Book",
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)

	if err != nil {
		panic(err)
	}

	book.ID = int64(id)

	err = repository.UpdateBook(database.DbConnection, book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book

	id, err := strconv.Atoi(c.Param("id"))
	book.ID = int64(id)

	err = repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Book",
	})
}
