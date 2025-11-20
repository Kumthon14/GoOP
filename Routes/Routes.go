package routes

import (
	control "Go_OOP/Controller"
	controllerfactory "Go_OOP/ControllerFactory"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	var controller any
	var controllerFactory *controllerfactory.ControllerFactory

	controller = nil
	controllerFactory = controllerFactory.GetIntance()

	authController := controllerFactory.GetController("AuthController").(*control.AuthController)

	grpNontoken := r.Group("auth-api")
	{
		grpNontoken.POST("login", authController.Login)
	}

	grpToken := r.Group("user-api", authController.IsAccess)
	{
		controller = controllerFactory.GetController("UserController").(*control.UserController)

		grpToken.POST("adduser", controller.CreateUser)
		grpToken.GET("user", controller.GetUsers)
		grpToken.GET("user/:id", controller.GetUserById)
		grpToken.PUT("user/:id", controller.UpdateUser)
		grpToken.DELETE("user/:id", controller.DeleteUser)
	}

	grpUpload := r.Group("/upload-api", authController.IsAccess)
	{
		controller = controllerFactory.GetController("FileManageController").(*control.FileManageController)

		grpUpload.GET("getUploadLists", controller.GetUploadLists)
		grpUpload.POST("uploadFile", controller.UploadFile)
	}

	grpDownload := r.Group("/downloadFile", authController.IsAccess)
	{
		controller = controllerFactory.GetController("FileManageController").(*control.FileManageController)

		grpDownload.GET("byId/:fileId", controller.DownloadFileById)
	}
}
