package author

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type authorRepositoryPostgres struct {
	db *gorm.DB
}

func NewAuthorRepositoryPostgres(db *gorm.DB) AuthorRepository {
	err := db.AutoMigrate(&Author{}).Error
	if err != nil {
		fmt.Println(err)
	}
	log.Print("Successfully connected to postgres in author service!")

	return &authorRepositoryPostgres{
		db: db,
	}
}

func (repo *authorRepositoryPostgres) createAuthor(author Author) (Author, error) {
	err := repo.db.Create(&author).Error
	if err != nil {
		return author, err
	}
	return author, err
}

func (repo *authorRepositoryPostgres) getAuthor(authorId uint) (Author, error) {
	author := Author{}
	err := repo.db.Where("id = ?", authorId).First(&author).Error
	if err != nil {
		return author, err
	}
	return author, err
}

func (repo *authorRepositoryPostgres) getAuthors() ([]Author, error) {
	authors := []Author{}
	err := repo.db.Find(&authors).Error
	if err != nil {
		return authors, err
	}
	return authors, err
}

func (repo *authorRepositoryPostgres) deleteAuthor(authorId uint) error {
	err := repo.db.Where("id = ?", authorId).Delete(Author{}).Error
	if err != nil {
		return err
	}
	return err
}

func (repo *authorRepositoryPostgres) updateAuthor(author Author) error {
	err := repo.db.Save(&author).Error
	if err != nil {
		return err
	}
	return err
}
