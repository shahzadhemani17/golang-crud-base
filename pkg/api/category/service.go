package category

type CategoryService struct {
	categoryRepository CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) CategoryService {
	return CategoryService{categoryRepository: categoryRepo}
}

func (svc *CategoryService) CreateCategory(category Category) (Category, error) {
	return svc.categoryRepository.createCategory(category)
}

func (svc *CategoryService) GetCategory(categoryId uint) (Category, error) {
	return svc.categoryRepository.getCategory(categoryId)
}

func (svc *CategoryService) GetCategories() ([]Category, error) {
	return svc.categoryRepository.getCategories()
}

func (svc *CategoryService) DeleteCategory(categoryId uint) error {
	return svc.categoryRepository.deleteCategory(categoryId)
}

func (svc *CategoryService) UpdateCategory(category Category) error {
	return svc.categoryRepository.updateCategory(category)
}
