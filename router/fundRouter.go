package router

import (
	"fund/controller/fund"
	"github.com/gin-gonic/gin"
)

func FundRouter(router *gin.RouterGroup)  {
	fundRouter := router.Group("api/fund")
	{
		fundRouter.POST("",fund.AddFunds)
		fundRouter.GET("",fund.GetFundDetail)
	}
}
