package main

import (
	routes "Go_OOP/Routes"
	_ "strings"
	_ "text/template"

	_ "github.com/gin-gonic/gin"
)

var err error

// @title Go Hello Api
// @version 1.0
// @destination Go Learning Project
// @termOfService http://agilerap.com/

// @contact.name API Support
// @contact.url http://agilerap.com
// @contact.email support@agilerap.com

// @license.name Agilerap
// @license.url http://agilerap.com/

// @schemas https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := routes.SetUpRouter()

	r.Run()
}
