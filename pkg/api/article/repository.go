package article

// AdRepository Used to store and retrieve screens
type ArticleRepository interface {
	createArticle(article Article) (Article, error)
	getArticle(articleId uint) (Article, error)
	getArticles() ([]Article, error)
	deleteArticle(articleId uint) error
	updateArticle(article Article) error
}
