package main

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/matryer/try.v1"
	"os"
	"time"
	"weight-tracker/pkg/api/article"
	"weight-tracker/pkg/api/author"
	"weight-tracker/pkg/api/category"
	"weight-tracker/pkg/repository"
)

const (
	DB_SERVICE_DIALECT = "DB_SERVICE_DIALECT"
	CATEGORY_FILE_NAME = "CATEGORY_FILE_NAME"
	//Filename           = "categories_data.json"
)

func main() {
	connectionString := "postgres://pg:pass@localhost:5432/crud?sslmode=disable"

	// setup database connection
	db, err := setupDatabase(connectionString)
	if err != nil {
		panic(err)
	}
	// create storage dependency

	err = repository.RunMigrations(connectionString)

	if err != nil {
		panic(err)
	}
	// create router dependency
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// create article service
	categoryRepo := category.NewCategoryRepositoryInFile(os.Getenv(CATEGORY_FILE_NAME))
	categoryService := category.NewCategoryService(categoryRepo)
	category.Routes(router, categoryService)

	// create article service
	authorRepo := author.NewAuthorRepositoryPostgres(db)
	authorService := author.NewAuthorService(authorRepo)
	author.Routes(router, authorService)

	// create article service
	articleRepo := article.NewArticleRepositoryInMem()
	articleService := article.NewArticleService(articleRepo)
	article.Routes(router, articleService, authorService, categoryService)

	// start the server
	API_SERVER_PORT := os.Getenv("PORT")
	if len(API_SERVER_PORT) == 0 {
		API_SERVER_PORT = "3000"
	}
	router.Run(fmt.Sprintf(":%v", API_SERVER_PORT))

	if err != nil {
		panic(err)
	}
}

func setupDatabase(connString string) (*gorm.DB, error) {
	// change "postgres" for whatever supported database you want to use
	var db *gorm.DB
	const attempts = 5
	err := try.Do(func(attempt int) (bool, error) {
		fmt.Printf("Connecting to db, attempt %v\n", attempt)
		var err error
		db, err = gorm.Open(os.Getenv(DB_SERVICE_DIALECT), connString)
		fmt.Printf("failed to connect database, attempt # %v", attempt)
		fmt.Println(fmt.Errorf("error: %w", err))
		if err == nil {
			return true, nil
		}
		sleepTime := attempt * attempts * 2
		time.Sleep(time.Second * time.Duration(sleepTime))
		return attempt < attempts, err
	})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	// ping the DB to ensure that it is connected
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
