package article

type ArticleService struct {
	articleRepository ArticleRepository
}

func NewArticleService(articleRepo ArticleRepository) ArticleService {
	return ArticleService{articleRepository: articleRepo}
}

func (svc *ArticleService) CreateArticle(article Article) (Article, error) {
	return svc.articleRepository.createArticle(article)
}

func (svc *ArticleService) GetArticle(articleId uint) (Article, error) {
	return svc.articleRepository.getArticle(articleId)
}

func (svc *ArticleService) GetArticles() ([]Article, error) {
	return svc.articleRepository.getArticles()
}

func (svc *ArticleService) DeleteArticle(articleId uint) error {
	return svc.articleRepository.deleteArticle(articleId)
}

func (svc *ArticleService) UpdateArticle(article Article) error {
	return svc.articleRepository.updateArticle(article)
}
