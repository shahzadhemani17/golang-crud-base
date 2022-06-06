package article

import (
	"errors"
)

type ArticleRepositoryInMemory struct {
	articles map[uint]Article
}

func NewArticleRepositoryInMem() ArticleRepository {
	return &ArticleRepositoryInMemory{
		articles: make(map[uint]Article),
	}
}

func (repo *ArticleRepositoryInMemory) createArticle(article Article) (Article, error) {
	article.Id = uint(len(repo.articles) + 1)
	repo.articles[article.Id] = article
	return article, nil
}

func (repo *ArticleRepositoryInMemory) getArticle(articleId uint) (Article, error) {

	article, ok := repo.articles[articleId]
	if !ok {
		return article, errors.New("article note found")
	}

	return article, nil
}

func (repo *ArticleRepositoryInMemory) getArticles() ([]Article, error) {
	articles := make([]Article, 0)
	for _, val := range repo.articles {
		articles = append(articles, val)
	}
	return articles, nil
}

func (repo *ArticleRepositoryInMemory) deleteArticle(articleId uint) error {
	delete(repo.articles, articleId)
	return nil
}

func (repo *ArticleRepositoryInMemory) updateArticle(article Article) error {
	repo.articles[article.Id] = article
	return nil
}
