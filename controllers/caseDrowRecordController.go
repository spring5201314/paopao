package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
)

func GetAllCaseDrowRecord(c *gin.Context) {
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	channel, err := models.GetCaseDrowRecord(n)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取提现记录成功",
		Data: gin.H{
			"data": channel,
		},
	})
	return
}

type oneCaseDrowRecordParams struct {
	id  string `form:"id" json:"id"`
	Page string `form:"page" json:"page"`

}

func DeleteCaseDrowRecord(c *gin.Context) {
	deleteParams := oneCaseDrowRecordParams{}
	c.Bind(&deleteParams)
	page, err := strconv.Atoi(deleteParams.Page)
	res, err := models.DeleteCaseDrowRecord(page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除体现记录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除体现记录成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

//定义修改结构体需要参数
type UpdateCaseDrowRecordParams struct {
	Id      string `form:"id" json:"id"`
	Context string `form:"text" json:"text"`
}

func UpdateCaseDrowRecord(c *gin.Context) {
	updateParams := UpdateCaseDrowRecordParams{}
	c.Bind(&updateParams)
	id, err := strconv.Atoi(updateParams.Id)
	res, err := models.UpdateCaseDrowRecord(id, updateParams.Context)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "修改取款记录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "修改取款记录成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}


