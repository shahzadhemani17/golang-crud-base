package author

type AuthorService struct {
	articleRepository AuthorRepository
}

func NewAuthorService(articleRepo AuthorRepository) AuthorService {
	return AuthorService{articleRepository: articleRepo}
}

func (svc *AuthorService) CreateAuthor(article Author) (Author, error) {
	return svc.articleRepository.createAuthor(article)
}

func (svc *AuthorService) GetAuthor(articleId uint) (Author, error) {
	return svc.articleRepository.getAuthor(articleId)
}

func (svc *AuthorService) GetAuthors() ([]Author, error) {
	return svc.articleRepository.getAuthors()
}

func (svc *AuthorService) DeleteAuthor(articleId uint) error {
	return svc.articleRepository.deleteAuthor(articleId)
}

func (svc *AuthorService) UpdateAuthor(article Author) error {
	return svc.articleRepository.updateAuthor(article)
}
