package main

import (
	"database/sql"
	"fmt"
	"log"
	"mini-project-quiz-3/controllers"
	"mini-project-quiz-3/database"
	"mini-project-quiz-3/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatalf("Error load env")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// router GIN
	router := gin.Default()

	// Bangun Datar
	// router.GET("/bangun-datar/:jenis", controllers.GetBangunDatar)

	router.GET("/bangun-datar/segitiga-sama-sisi", controllers.GetHasilSegitiga)
	router.GET("/bangun-datar/persegi", controllers.GetHasilPersegi)
	router.GET("/bangun-datar/persegi-panjang", controllers.GetHasilPersegiPanjang)
	router.GET("/bangun-datar/lingkaran", controllers.GetHasilLingkaran)

	// gin basic auth middleware ==========
	// authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"admin":  "password",
	// 	"editor": "secret",
	// }))
	// ====================================

	authorized := router.Group("/")
	authorized.Use(middleware.Auth())

	// Category
	router.GET("/categories", controllers.GetCategory)
	router.GET("/categories/:id/books", controllers.GetBookByCategoryId)
	authorized.POST("/categories", controllers.InsertCategory)
	authorized.PUT("/categories/:id", controllers.UpdateCategory)
	authorized.DELETE("/categories/:id", controllers.DeleteCategory)

	// Book
	router.GET("/books", controllers.GetBook)
	authorized.POST("/books", controllers.InsertBook)
	authorized.PUT("/books/:id", controllers.UpdateBook)
	authorized.DELETE("/books/:id", controllers.DeleteBook)
	router.Run(":" + os.Getenv("PORT"))
}
