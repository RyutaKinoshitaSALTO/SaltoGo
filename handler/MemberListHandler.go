package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MemberListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Members)
}
