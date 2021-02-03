package main

import (
	middleware "fund/middleware/cors"
	routerInit "fund/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main()  {
	gin.SetMode(viper.GetString("gin.mod"))
	router := gin.Default()

	//	解决跨域问题
	router.Use(middleware.Cors())

	//日期格式化
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	//	接口api
	apiRouter := router.Group("/v1")
	routerInit.InitRouter(apiRouter)

	router.Run(viper.GetString("port"))
}
