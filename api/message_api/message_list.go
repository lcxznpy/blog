package message_api

import (
	"blog_server/global"
	"blog_server/models"
	"blog_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListView(c *gin.Context) {
	_cliams, _ := c.Get("claims") //set的东西是个泛型
	claims := _cliams.(*jwts.CustomClaims)
	var messList []models.MessageModel
	global.DB.First(&messList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	fmt.Println(messList)
}
