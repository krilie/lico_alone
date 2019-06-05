package user

import (
	"github.com/gin-gonic/gin"
	"github.com/krilie/lico_alone/application/user"
	"github.com/krilie/lico_alone/control/middleware"
)

var appUser user.AppUser

type UserCtrl struct{}

type UserCtrler interface {
	Login(c *gin.Context)
	GetInfo(c *gin.Context)
	Logout(c *gin.Context)
	ValidClientAccToken(c *gin.Context)
	Valid(c *gin.Context)
	Register(c *gin.Context)
	UserCtrlAccounter
}

type UserCtrlAccounter interface {
	DeleteBill(c *gin.Context)
	GetAccountHistory(c *gin.Context)
	GetAccountInfo(c *gin.Context)
	AddBill(c *gin.Context)
	AddAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
	GetMonthSummary(c *gin.Context)
	GetTimeZoneSummary(c *gin.Context)
}

func Init(group *gin.RouterGroup) {
	var userCtrl UserCtrl
	//用户基础
	userBase := group.Group("/user")
	userBase.POST("/login", userCtrl.Login)
	userBase.POST("/logout", middleware.CheckAuthToken(), userCtrl.Logout) // 要token有效
	userBase.POST("/register", userCtrl.Register)
	userBase.GET("/valid", userCtrl.Valid)                                // 不要登录，要有客户端的key
	userBase.GET("/valid_client_acc_token", userCtrl.ValidClientAccToken) //不要权限的
	// account
	accountGp := group.Group("/account")
	accountGp.Use(middleware.CheckAuthToken())
	accountGp.POST("/delete_bill", userCtrl.DeleteBill)
	accountGp.GET("/account_history", userCtrl.GetAccountHistory)
	accountGp.GET("/account_info", userCtrl.GetAccountInfo)
	accountGp.POST("/add_bill", userCtrl.AddBill)
	accountGp.POST("/add_account", userCtrl.AddAccount)
	accountGp.POST("/delete_account", userCtrl.DeleteAccount)
	accountGp.GET("/month_summary", userCtrl.GetMonthSummary)
	accountGp.GET("/time_zone_summary", userCtrl.GetTimeZoneSummary)
}
