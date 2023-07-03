package types

type Todo struct {
	ID     uint   `gorm:"primary key;autoIncrement" json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}
