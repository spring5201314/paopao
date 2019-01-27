package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"paopao/models"
	"strconv"
)

func GetAllBankCard(c *gin.Context) {
	page := c.DefaultPostForm("page", "0")
	n, err := strconv.Atoi(page)
	channel, err := models.GetAllBankCard(n)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "登录失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "获取银行卡成功",
		Data: gin.H{
			"data": channel,
		},
	})
	return
}

type oneBankCardParams struct {
	id  string `form:"id" json:"id"`
	Page string `form:"page" json:"page"`

}

func DeleteBankCard(c *gin.Context) {
	deleteParams := oneBankCardParams{}
	c.Bind(&deleteParams)
	page, err := strconv.Atoi(deleteParams.Page)
	res, err := models.DeleteBankCard(page)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "删除银行卡失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "删除银行卡成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}

//定义修改结构体需要参数
type UpdateBankCardParams struct {
	Id      string `form:"id" json:"id"`
	Context string `form:"text" json:"text"`
}

func UpdateBankCard(c *gin.Context) {
	updateParams := UpdateChannelParams{}
	c.Bind(&updateParams)
	id, err := strconv.Atoi(updateParams.Id)
	res, err := models.UpdateBankCard(id, updateParams.Context)
	if err != nil {
		c.JSON(http.StatusOK, ApiRes{
			Code: 1,
			Msg:  "修改银行卡失败",
		})
		return
	}
	c.JSON(http.StatusOK, ApiRes{
		Code: 0,
		Msg:  "修改银行卡成功",
		Data: gin.H{
			"data": res,
		},
	})
	return
}



func CreatBankCard(c *gin.Context) {

	regParams := models.T_member_bank{}
	err := c.BindJSON(&regParams)
	c.Bind(&regParams)

	//userAccount, bankAccountName, bankCode,bankCardNo,bankAddress,createTime string,traceId int,isDel,isDefault boo
	channel, err := models.AddBankCard(regParams.User_account, regParams.Bank_account_name,regParams.Bank_code,regParams.Bank_card_no,regParams.Bank_address,regParams.Creation_time,regParams.Trace_id,regParams.Is_del,regParams.Is_default)
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