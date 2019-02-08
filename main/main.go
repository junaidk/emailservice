package main

import (
	"emailservice/constants"
	"emailservice/controllers"
	"emailservice/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.InitFlags()
	fmt.Println(constants.Username)
	fmt.Println(constants.Password)
	fmt.Println(constants.FromAddress)

	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	ctrl := controllers.NewController()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/send", ctrl.SendEmail)
	}

	router.Run(":9300") // listen and serve on 0.0.0.0:8080

}
