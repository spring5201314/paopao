package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
)

func GetAllChannel(c *gin.Context) {
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	channel, err := models.GetAllChannel(n)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取通道成功",
		Data: gin.H{
			"data": channel,
		},
	})
	return
}

type oneChannelParams struct {
	id  string `form:"id" json:"id"`
	Page string `form:"page" json:"page"`

}

func DeleteChannel(c *gin.Context) {
	deleteParams := oneChannelParams{}
	c.Bind(&deleteParams)
	page, err := strconv.Atoi(deleteParams.Page)
	res, err := models.DeleteChannel(page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除通道失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除通道成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

//定义修改结构体需要参数
type UpdateChannelParams struct {
	Id      string `form:"id" json:"id"`
	Context string `form:"text" json:"text"`
}

func UpdateChannel(c *gin.Context) {
	updateParams := UpdateChannelParams{}
	c.Bind(&updateParams)
	id, err := strconv.Atoi(updateParams.Id)
	res, err := models.UpdateChannel(id, updateParams.Context)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "修改通道失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "修改通道成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}



func CreatChannel(c *gin.Context) {

	regParams := models.T_channel{}
	err := c.BindJSON(&regParams)
	c.Bind(&regParams)

	channel, err := models.AddChannel(regParams.Channel_code, regParams.Channel_name,regParams.Channel_status,regParams.Rate)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "创建通道成功",
		Data: gin.H{
			"data": channel,
		},
	})
	return
}