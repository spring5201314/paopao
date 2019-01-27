

package controllers

import (
	"fmt"
	"paopao/config"
	"paopao/models"
	"paopao/util"

	"github.com/gin-gonic/gin"
	"net/http"

	"time"
	// "io/ioutil"
)

func Register(c *gin.Context) {
	regParams := models.T_merchants_user{}
	err := c.BindJSON(&regParams)
	fmt.Println(regParams, "--------------------------")
	if regParams.Pass_word == "" || regParams.User_name == "" {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "姓名或密码不能为空",
			})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "获取数据错误",
			})
		return
	}
	if len(regParams.User_name) < 0 || len(regParams.Pass_word) < 0 {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "账号或密码不能为空",
			})
		return
	}
	hadUser := models.GetName(regParams.User_name) //判断是否已经注册
	if hadUser {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "用户已经存在",
			})
		return
	}
	//name, pass,tracekey,corporationName,idNumber,businessAddress,email,qq,regTime string,traceId,phoneNum,sex int
	user, err := models.UserRegister(regParams.User_name,regParams.Pass_word,regParams.Trace_key,regParams.Corporation_name,regParams.Id_number,regParams.Business_address,regParams.Email,regParams.Qq,regParams.RegTime,regParams.Trace_id,regParams.Phone_number,regParams.Sex)
	if err != nil {
		c.JSON(http.StatusOK,
			ApiRes{
				Code: 1,
				Msg:  "登录数据格式不正确！",
			})
		return
	}
	token, err := utils.Encrypt(fmt.Sprintf("%d:%d", user.Id, time.Now().Unix()+config.TOKEN_EXPIRE_TIME), []byte(config.EncryptKey))
	c.JSON(http.StatusOK,
		ApiRes{
			Code: 0,
			Msg:  "成功注册",
			Data: gin.H{
				"token": token,
			},
		})
	return
}
