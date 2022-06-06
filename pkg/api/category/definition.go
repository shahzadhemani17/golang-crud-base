package category

//Category model
type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func RemoveIndex(c []Category, index int) []Category {
	return append(c[:index], c[index+1:]...)
}
