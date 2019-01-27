package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
)

func GetAllUser(c *gin.Context) {
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	channel, err := models.GetAllUser(n)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取用户成功",
		Data: gin.H{
			"data": channel,
		},
	})
	return
}

type oneUserParams struct {
	id  string `form:"id" json:"id"`
	Page string `form:"page" json:"page"`

}

func DeleteUser(c *gin.Context) {
	deleteParams := oneUserParams{}
	c.Bind(&deleteParams)
	page, err := strconv.Atoi(deleteParams.Page)
	res, err := models.DeleteUser(page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除用户成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

//定义修改结构体需要参数
type UpdateUserParams struct {
	Id      string `form:"id" json:"id"`
	Context string `form:"text" json:"text"`
}

func UpdateUser(c *gin.Context) {
	updateParams := UpdateUserParams{}
	c.Bind(&updateParams)
	id, err := strconv.Atoi(updateParams.Id)
	res, err := models.UpdateUser(id, updateParams.Context)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "修改用户失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "修改用户成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}



