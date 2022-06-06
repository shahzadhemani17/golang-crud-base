package category

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type categoryRepositoryInFile struct {
	file *os.File
}

func NewCategoryRepositoryInFile(filename string) CategoryRepository {
	myFile, e := os.Create(filename)
	if e != nil {
		log.Fatal(e)
	}
	return &categoryRepositoryInFile{
		file: myFile,
	}
}

func (repo *categoryRepositoryInFile) createCategory(category Category) (Category, error) {

	jsonFileData, err := repo.getCategories()
	category.Id = uint(len(jsonFileData) + 1)

	jsonFileData = append(jsonFileData, category)
	file, err := json.MarshalIndent(jsonFileData, "", " ")

	err = ioutil.WriteFile(repo.file.Name(), file, 0644)
	if err != nil {
		return category, err
	}
	return category, err
}

func (repo *categoryRepositoryInFile) getCategory(categoryId uint) (Category, error) {
	categories, err := repo.getCategories()
	if err != nil {
		return Category{}, err
	}

	for _, c := range categories {
		if c.Id == categoryId {
			return c, nil
		}
	}
	return Category{}, errors.New("category not found")
}

func (repo *categoryRepositoryInFile) getCategories() ([]Category, error) {
	jsonFileData := []Category{}
	data, err := ioutil.ReadFile(repo.file.Name())
	if err != nil {
		return jsonFileData, err
	}
	err = json.Unmarshal(data, &jsonFileData)
	if err != nil {
		return jsonFileData, err
	}
	return jsonFileData, err
}

func (repo *categoryRepositoryInFile) deleteCategory(categoryId uint) error {
	categories, err := repo.getCategories()
	if err != nil {
		return err
	}

	for i, c := range categories {
		if c.Id == categoryId {
			RemoveIndex(categories, i)
		}
	}

	byteData, err := json.Marshal(categories)

	err = ioutil.WriteFile(repo.file.Name(), byteData, 0644)
	if err != nil {
		return err
	}
	return err
}

func (repo *categoryRepositoryInFile) updateCategory(category Category) error {
	categories, err := repo.getCategories()
	if err != nil {
		return err
	}

	for i, c := range categories {
		if c.Id == category.Id {
			categories[i] = category
		}
	}

	byteData, err := json.MarshalIndent(categories, "", " ")

	err = ioutil.WriteFile(repo.file.Name(), byteData, 0644)
	if err != nil {
		return err
	}
	return err
}
