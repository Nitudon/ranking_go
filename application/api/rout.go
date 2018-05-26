package api

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	router.GET("/regist", RegistUser)
	router.GET("/ranking", GetRanking)
}
