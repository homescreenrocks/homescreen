package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/homescreenrocks/homescreen/core/backend/modulemanager"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	m := gin.Default()

	var execMode = flag.Bool("exec", false, "spawns plugins during startup")
	flag.Parse()
	mm := modulemanager.New(*execMode)

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// every file in frontend will be served

	log.Print(pwd + "/core/frontend/index.html")
	m.StaticFile("/", pwd+"/core/frontend/index.html")
	m.Static("/www", pwd+"/core/frontend")

	rootGroup := m.Group("/api/v1")
	{
		mm.RegisterRouterGroup(rootGroup.Group("/modules"))
	}

	m.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})
	mm.ScanForModules()
	log.Println("Server is running...")

	m.Run(":3000")

}
