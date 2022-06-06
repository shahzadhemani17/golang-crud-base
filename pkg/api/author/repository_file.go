package author

type authorRepositoryInFile struct {
	authors map[uint]Author
}

func NewAuthorRepositoryInFile() AuthorRepository {
	return &authorRepositoryInFile{
		authors: make(map[uint]Author),
	}
}

func (repo *authorRepositoryInFile) createAuthor(article Author) (Author, error) {
	return Author{}, nil
}

func (repo *authorRepositoryInFile) getAuthor(articleId uint) (Author, error) {
	return Author{}, nil
}

func (repo *authorRepositoryInFile) getAuthors() ([]Author, error) {
	return nil, nil
}

func (repo *authorRepositoryInFile) deleteAuthor(articleId uint) error {
	return nil
}

func (repo *authorRepositoryInFile) updateAuthor(article Author) error {
	return nil
}
