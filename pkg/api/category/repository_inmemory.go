package category

type categoryRepositoryInMemory struct {
	categories map[uint]Category
}

func NewCategoryRepositoryInMem() CategoryRepository {
	return &categoryRepositoryInMemory{
		categories: make(map[uint]Category),
	}
}

func (repo *categoryRepositoryInMemory) createCategory(article Category) (Category, error) {
	return Category{}, nil
}

func (repo *categoryRepositoryInMemory) getCategory(articleId uint) (Category, error) {
	return Category{}, nil
}

func (repo *categoryRepositoryInMemory) getCategories() ([]Category, error) {
	return nil, nil
}

func (repo *categoryRepositoryInMemory) deleteCategory(articleId uint) error {
	return nil
}

func (repo *categoryRepositoryInMemory) updateCategory(article Category) error {
	return nil
}
