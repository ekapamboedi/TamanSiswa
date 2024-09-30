package main

import (
	"github.com/gin-gonic/gin"

	"TamanSiswa/config"
	"TamanSiswa/middleware"
	"TamanSiswa/model"
	menuaccess "TamanSiswa/services/menuAccess"
	"TamanSiswa/services/salary"
	"TamanSiswa/services/auth"
	
	// "TamanSiswa/services/auth/response"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	
	config.Init()
	// config.InitOss()
	model.Init()

	if config.ENVIRONMENT == "Production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	// Maximum Size for uploaded file
	r.MaxMultipartMemory = 2 << 20

	r.GET("", func(ctx *gin.Context) {
		response := map[string]string{}
		response["status"] = "Success"
		response["message"] = "Hello, Welcome to Taman Siswa"
		ctx.IndentedJSON(200, response)
	})

	authRoute := r.Group("/api/v1/login", middleware.AuthenticateOnly)
	{
		var res response.Response
		// auth.AuthRoute(authRoute)
	
		authRoute.GET("/protected", func(ctx *gin.Context) {
			
			res.Status = "Login succed"
			res.Message = "Success"
			ctx.JSON(IndentedJSON, res)
		})
	}

	menuAccessRoute := r.Group("api/v1/menu", middleware.AuthenticateOnly, middleware.AdminAuthenticate)
	{
		menuaccess.MenuRoute(menuAccessRoute)
	}

	salaryRoute := r.Group("api/v1/salaries", middleware.Authenticate)
	{
		salary.SalaryRoute(salaryRoute)
	}

	r.Run(":" + config.SERVER_PORT)
}
