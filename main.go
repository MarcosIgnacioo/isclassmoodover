package main

import (
	// "log"
	// "os"
	// "path/filepath"

	// "github.com/MarcosIgnacioo/classmoodls/controllers"
	// "github.com/gin-gonic/gin"
	// "github.com/playwright-community/playwright-go"

	pw "github.com/MarcosIgnacioo/classmoodls/playwright"
	types "github.com/MarcosIgnacioo/classmoodls/types"
)

func main() {
	user := types.User{Username: "marcosignc_21", Password: "sopitasprecio"}
	pw.StartScrapping(user)
	// state := os.Getenv("STATE")
	// var dir string
	// var err error
	// if state == "prod" {
	// 	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	// } else {
	// 	dir = "."
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// r := gin.Default()
	// r.LoadHTMLGlob(dir + "/public/templates/*")
	// r.Static("/assets", dir+"/assets")
	// r.GET("/", controllers.LogIn)
	// r.GET("/wa", controllers.Test)
	// r.POST("/LogIn", controllers.LogInPost)
	// var port string
	// if state == "dev" {
	// 	port = "3030"
	// } else {
	// 	port = "3000"
	// }
	// r.Run("0.0.0.0:" + port)
}
