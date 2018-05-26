package domain

type User struct {
	ID    string `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Score int    `db:"score" json:"score"`
}
