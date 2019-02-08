package main

import (
	"emailservice/controllers"
	"emailservice/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.InitFlags()

	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	ctrl := controllers.NewController()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/send", ctrl.SendEmail)
	}

	router.Run(":80") // listen and serve on 0.0.0.0:8080

}
