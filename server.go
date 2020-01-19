package main

import (
	"github.com/gin-gonic/gin"
	"go-disk/handler"
	"log"
)

func main() {
	router := gin.Default()

	filesGroup := router.Group("/files")
	usersGroup := router.Group("/users")

	DispatchHandlerGroup(filesGroup, usersGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func DispatchHandlerGroup(rgs ...*gin.RouterGroup) {
	for _, rg := range rgs {
		switch rg.BasePath() {
		case "/files":
			handler.FilesServiceHandler{BashPath: rg.BasePath()}.Init(rg)
		case "/users":
			handler.UserServiceHandler{BashPath: rg.BasePath()}.Init(rg)
		default:
			log.Printf("error handler group: %s", rg.BasePath())
		}
	}

}