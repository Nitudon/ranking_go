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

	infra.DB.First("id = ?", id, &user)
	user.Score, _ = strconv.Atoi(score)
	infra.DB.Save(&user)

	ranking := []*domain.User{}

	infra.DB.Order("score desc").Find(&ranking)

	g.JSON(
		200,
		ranking,
	)

}
