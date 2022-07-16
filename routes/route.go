package routes

import (
	"SaltoGo/handler"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/member", handler.MemberHandler)
	router.GET("/member/list", handler.MemberListHandler)
	router.PUT("/member/update/:id", handler.UpdateMemberHandler)
	router.DELETE("/member/delete/:id", handler.DeleteMemberHandler)
	return router
}
