package author

type AuthorRepositoryInMemory struct {
	authors map[uint]Author
}

func NewAuthorRepositoryInMem() AuthorRepository {
	return &AuthorRepositoryInMemory{
		authors: make(map[uint]Author),
	}
}

func (repo *AuthorRepositoryInMemory) createAuthor(article Author) (Author, error) {
	return Author{}, nil
}

func (repo *AuthorRepositoryInMemory) getAuthor(articleId uint) (Author, error) {
	return Author{}, nil
}

func (repo *AuthorRepositoryInMemory) getAuthors() ([]Author, error) {
	return nil, nil
}

func (repo *AuthorRepositoryInMemory) deleteAuthor(articleId uint) error {
	return nil
}

func (repo *AuthorRepositoryInMemory) updateAuthor(article Author) error {
	return nil
}
