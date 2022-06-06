package article

//Article model
type Article struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryId  uint   `json:"category_id"`
	AuthorId    uint   `json:"author_id"`
}

type ArticleResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Author      string `json:"author"`
}
