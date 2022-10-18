package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//ginSwagger.WrapHandler(swaggerFiles.Handler,
	//	ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	//	ginSwagger.DefaultModelsExpandDepth(-1))

	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("usr/createUser", service.CreateUser)

	return r

}
