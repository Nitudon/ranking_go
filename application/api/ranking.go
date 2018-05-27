package api

import (
	domain "../../domain"
	infra "../../infra"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func RegistUser(g *gin.Context) {
	name := g.Query("name")

	id := xid.New().String()

	user := new(domain.User)
	user.ID = id
	user.Name = name
	user.Score = -1

	infra.DB.Create(&user)

	g.String(
		200,
		id,
	)
}

func GetRanking(g *gin.Context) {
	id := g.Query("id")
	score := g.Query("score")

	user := new(domain.User)
	var rank int

	infra.DB.First("id = ?", id, &user)
	user.Score, _ = strconv.Atoi(score)
	infra.DB.Where("score > ?", user.Score).Count(&rank)
	user.Rank = rank + 1
	infra.DB.Save(&user)

	ranking := []*domain.User{}

	infra.DB.Where("score >= ?", 0).Order("score asc").Find(&ranking)

	g.JSON(
		200,
		ranking,
	)

}
