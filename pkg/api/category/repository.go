package category

// CategoryRepository Used to store and retrieve screens
type CategoryRepository interface {
	createCategory(article Category) (Category, error)
	getCategory(articleId uint) (Category, error)
	getCategories() ([]Category, error)
	deleteCategory(articleId uint) error
	updateCategory(article Category) error
}
