// Package routers provides ...
package routers

import (
	"github.com/gin-gonic/gin"
	"paopao/controllers"
	"paopao/middleware"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.LoadHTMLGlob("templates/*")
	// router.Static("/assets", "./assets")
	router.Use(middleware.Cors())
	router.Use(middleware.AllUrlPath)
	router.POST("/Register", controllers.Register)                         // 注册
	router.POST("/LoginIn", controllers.LoginIn) //登录

	//通道模块
	router.POST("/getAllChannel", controllers.GetAllChannel)
	router.POST("/deleteChannel", controllers.DeleteChannel)
	router.POST("/createChannel", controllers.CreatChannel)
	router.POST("/updateChannel", controllers.UpdateChannel)
    //银行卡模块
	router.POST("/getAllBankCard", controllers.GetAllBankCard)
	router.POST("/addBankCard", controllers.CreatBankCard)
	router.POST("/deleteBankCard", controllers.DeleteBankCard)
	router.POST("/updateBankCard", controllers.UpdateBankCard)

	//取款记录模块
	router.POST("/getAllCaseDrowRecord", controllers.GetAllCaseDrowRecord)
	router.POST("/deleteCaseDrowRecord", controllers.DeleteCaseDrowRecord)
	router.POST("/updateCaseDrowRecord", controllers.UpdateCaseDrowRecord)

	//存款记录模块
	router.POST("/getAllCaseFlowRecord", controllers.GetAllCaseFlowRecord)
	router.POST("/deleteCaseFlowDrowRecord", controllers.DeleteCaseFlowDrowRecord)
	router.POST("/updateCaseFlowRecord", controllers.UpdateCaseFlowRecord)


	jwtrouter := router.Group("/jwt", middleware.UserAuth()) //token
	// jwtrouter := router.Group("/jwt") //token
	{
		//获取用户信息
		jwtrouter.POST("/GetUser", controllers.GetUser)
		//用户其它路由
		jwtrouter.POST("/getAllUser", controllers.GetAllUser)
		jwtrouter.POST("/updateUser", controllers.UpdateUser)
		jwtrouter.POST("/deleteUser", controllers.DeleteUser)

		//upload
		jwtrouter.POST("/Upload", controllers.Upload)

	}
	return router
}
