package author

// AuthorRepository Used to store and retrieve screens
type AuthorRepository interface {
	createAuthor(article Author) (Author, error)
	getAuthor(articleId uint) (Author, error)
	getAuthors() ([]Author, error)
	deleteAuthor(articleId uint) error
	updateAuthor(article Author) error
}
