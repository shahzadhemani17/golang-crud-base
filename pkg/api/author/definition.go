package author

//Article model
type Author struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Email string `json:"email"`
}
