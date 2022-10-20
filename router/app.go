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
	r.GET("/usr/getUserList", service.GetUserList)
	r.GET("usr/createUser", service.CreateUser)
	r.GET("usr/deleteUser", service.DeleteUser)
	r.POST("usr/updateUser", service.UpdateUser)
	r.POST("usr/findUserByNameAndPwd", service.FindUserByNameAndPwd)

	//发送消息
	r.GET("usr/sendMsg", service.SengMsg)

	return r

}
