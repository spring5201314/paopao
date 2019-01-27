package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
)

func GetAllCaseFlowRecord(c *gin.Context) {
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	channel, err := models.GetCaseFlowRecord(n)
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

type oneCaseFlowRecordParams struct {
	id  string `form:"id" json:"id"`
	Page string `form:"page" json:"page"`

}

func DeleteCaseFlowDrowRecord(c *gin.Context) {
	deleteParams := oneCaseFlowRecordParams{}
	c.Bind(&deleteParams)
	page, err := strconv.Atoi(deleteParams.Page)
	res, err := models.DeleteCaseFlowRecord(page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除存款记录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除存款记录成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

//定义修改结构体需要参数
type UpdateCaseFlowRecordParams struct {
	Id      string `form:"id" json:"id"`
	Context string `form:"text" json:"text"`
}

func UpdateCaseFlowRecord(c *gin.Context) {
	updateParams := UpdateCaseFlowRecordParams{}
	c.Bind(&updateParams)
	id, err := strconv.Atoi(updateParams.Id)
	res, err := models.UpdateCaseFlowRecord(id, updateParams.Context)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "修改存款记录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "修改存款记录成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}


