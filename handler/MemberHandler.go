package handler

import (
	"SaltoGo/model"
	"encoding/json"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var Members []model.Member

func init() {
	Members = make([]model.Member, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &Members)
}

func MemberHandler(c *gin.Context) {
	var Member model.Member
	// リクエストデータ取り出す
	if err := c.ShouldBindJSON(&Member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	// ID作成
	Member.ID = xid.New().String()
	// 現在時刻を追加
	Member.PublishedAt = time.Now()
	Members = append(Members, Member)
	c.JSON(http.StatusOK, Member)
}
