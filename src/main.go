package main

import (
	_ "github.com/go-sql-driver/mysql"
	"app"
	"controllers"
)

func main() {
	application := app.New()
	application.Get("site", &controllers.SiteController{})
	application.Run(":8080")
}
