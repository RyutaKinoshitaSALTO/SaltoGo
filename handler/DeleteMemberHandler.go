package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteMemberHandler(c *gin.Context) {
	id := c.Param("id")

	memberIndex := -1
	for i := 0; i < len(Members); i++ {
		if Members[i].ID == id {
			memberIndex = i
		}
	}

	if memberIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "recipe not found"})
		return
	}

	Members = append(Members[:memberIndex], Members[memberIndex+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted"})
}
