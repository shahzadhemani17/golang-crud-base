package category

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type categoryRepositoryPostgres struct {
	db *gorm.DB
}

func NewCategoryRepositoryPostgres(db *gorm.DB) CategoryRepository {
	err := db.AutoMigrate(&Category{}).Error
	if err != nil {
		fmt.Println(err)
	}
	log.Print("Successfully connected to postgres in category service!")

	return &categoryRepositoryPostgres{
		db: db,
	}
}

func (repo *categoryRepositoryPostgres) createCategory(category Category) (Category, error) {
	return Category{}, nil
}

func (repo *categoryRepositoryPostgres) getCategory(categoryId uint) (Category, error) {
	return Category{}, nil
}

func (repo *categoryRepositoryPostgres) getCategories() ([]Category, error) {
	return nil, nil
}

func (repo *categoryRepositoryPostgres) deleteCategory(categoryId uint) error {
	return nil
}

func (repo *categoryRepositoryPostgres) updateCategory(category Category) error {
	return nil
}
