package controllers

import (
	"mini-project-quiz-3/database"
	"mini-project-quiz-3/repository"
	"mini-project-quiz-3/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertCategory(c *gin.Context) {
	var category structs.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)

	if err != nil {
		panic(err)
	}

	category.ID = int64(id)

	err = repository.UpdateCategory(database.DbConnection, category)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCategory(c *gin.Context) {
	var category structs.Category

	id, err := strconv.Atoi(c.Param("id"))
	category.ID = int64(id)

	err = repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Category",
	})
}

func GetBookByCategoryId(c *gin.Context) {
	var (
		result gin.H
	)
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)

	category.ID = int64(id)

	books, err := repository.GetBookByCategory(database.DbConnection, category)

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
