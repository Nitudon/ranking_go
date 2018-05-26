package domain

type User struct {
	ID    string `db:"id" json:"ID"`
	Rank  int    `db:"rank" json:"Rank"`
	Name  string `db:"name" json:"Name"`
	Score int    `db:"score" json:"Score"`
}
