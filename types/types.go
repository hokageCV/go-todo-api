package types

type Todo struct {
	ID     string `json:"id" db:"primary_key"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
