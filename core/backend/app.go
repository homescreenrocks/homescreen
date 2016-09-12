package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/homescreenrocks/homescreen/core/backend/modulemanager"
	"github.com/homescreenrocks/homescreen/core/backend/routes"
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

	r := routes.New(&mm)

	rootGroup := m.Group("/api/v1")
	{
		rootGroup.GET("/getmodules", r.GetModules)

		rootGroup.GET("/getmodulesasjson", r.GetModulesAsJSON)
		rootGroup.GET("/findmodules", r.FindModules)

		rootGroup.POST("/registermodule", r.RegisterModule)

		moduleGroup := rootGroup.Group("/modules/:module")
		{
			moduleGroup.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "Parameter is %s", c.Param("module"))
			})
			moduleGroup.GET("/execute", r.ExecuteModules)
		}

		settingsGroup := rootGroup.Group("/settings")
		{
			settingsGroup.GET("/", r.GetSettings)
			settingsGroup.GET("/:module", r.SetSettingsPerModule)
			settingsGroup.POST("/:module/:key", r.SetSettingsPerModulePerKey)
		}
	}

	m.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})
	mm.ScanForModules()
	log.Println("Server is running...")

	m.Run(":3000")

}
