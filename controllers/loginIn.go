

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/config"
	"paopao/models"
	"paopao/util"

	// "strconv"
	"time"
)

func GetUser(c *gin.Context) {
	uinfoIr, _ := c.Get("udata")
	uidIr, _ := c.Get("uid")
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", uidIr, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "生成token失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "登录成功",
		Data: gin.H{
			"token": token,
			"user":  uinfoIr.(*models.T_merchants_user),
		},
	})
	return
}
func LoginIn(c *gin.Context) {
	loginParams := models.T_merchants_user{}
	err := c.Bind(&loginParams)
	if len(loginParams.User_name) <= 0 || len(loginParams.Pass_word) <= 0 {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "账号或密码不能为空",
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录数据格式不正确！",
		})
		return
	}
	user, err := models.UserLogin(loginParams.User_name, loginParams.Pass_word)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", user.Id, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "登录成功",
		Data: gin.H{
			"token": token,
			// "user":  &userData,
			"user": user,
		},
	})
	return
}
