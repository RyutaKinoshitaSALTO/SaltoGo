package handler

import (
	"SaltoGo/model"
	_ "SaltoGo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateMemberHandler(c *gin.Context) {
	id := c.Param("id")
	var member model.Member

	if err := c.ShouldBindJSON(&member); err != nil {
		// エラーの場合は、400エラーを返す
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	targetMemberIndex := -1
	for i := 0; i < len(Members); i++ {
		if Members[i].ID == id {
			targetMemberIndex = i
		}
	}

	if targetMemberIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found"})
		return
	}

	// 更新する
	Members[targetMemberIndex] = member
	c.JSON(http.StatusOK, member)
}