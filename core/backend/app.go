package main

import (
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

	mm := modulemanager.New()
	//m.Map(&mm)

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// every file in frontend will be served

	m.Use(apiMiddleware(&mm))
	log.Print(pwd + "/core/frontend/index.html")
	m.StaticFile("/", pwd+"/core/frontend/index.html")
	m.Static("/www", pwd+"/core/frontend")

	//m.GET("/", routes.Root)

	rootGroup := m.Group("/api/v1")
	{
		rootGroup.GET("/getmodules", routes.GetModules)

		rootGroup.GET("/getmodulesasjson", routes.GetModulesAsJSON)
		rootGroup.GET("/findmodules", routes.FindModules)

		moduleGroup := rootGroup.Group("/modules/:module")
		{
			moduleGroup.GET("/", func(c *gin.Context) {
				c.String(http.StatusOK, "Parameter is %s", c.Param("module"))
			})
			moduleGroup.GET("/execute", routes.ExecuteModules)
		}

		settingsGroup := rootGroup.Group("/settings")
		{
			settingsGroup.GET("/", routes.GetSettings)
			settingsGroup.GET("/:module", routes.SetSettingsPerModule)
			settingsGroup.POST("/:module/:key", routes.SetSettingsPerModulePerKey)
		}
	}

	//m.GET("/", routes.Root)

	/*
		m.NotFound(func(ctx *gin.Context) (int, string) {
			// Custom handle for 404
			return 404, "Could not find: " + ctx.Req.RequestURI
		})
	*/
	m.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})
	mm.ScanForModules()
	log.Println("Server is running...")
	//log.Println(mm.GetAllModules()["schnubbelmodule"])
	//log.Println(mm.GetAllModules()["weathermodule"])

	//log.Println("==================================")
	//log.Println(mm.GetAllModulesAsJSON())
	//log.Println(http.ListenAndServe("0.0.0.0:3000", m))
	m.Run(":3000")

}

func apiMiddleware(mm *modulemanager.ModuleManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mm", &mm)
		c.Next()
	}
}
