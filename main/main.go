package main

import (
	"emailservice/controllers"
	"emailservice/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func main() {
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	utils.InitFlags()

	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	ctrl := controllers.NewController()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/send", ctrl.SendEmail)
	}

	router.Run(":8080") // listen and serve on 0.0.0.0:8080

}
