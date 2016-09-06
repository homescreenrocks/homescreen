package routes

import (
	"log"
	"net/http"

	"github.com/homescreenrocks/homescreen/core/backend/modulemanager"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
	c.Redirect(301, "/index.html")
}

func GetModules(c *gin.Context) {
	mm := c.MustGet("mm").(*modulemanager.ModuleManager)
	c.JSON(http.StatusOK, (*mm).GetAllModules())
}

func GetModulesAsJSON(c *gin.Context) {
	mm := c.MustGet("mm").(*modulemanager.ModuleManager)
	c.JSON(http.StatusOK, (*mm).GetAllModules())
}

func FindModules(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI)
}

func ExecuteModules(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI + "var is: " + c.Param("module"))
}

func GetSettings(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI + " simple types")
}

func SetSettingsPerModule(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI + "var is: " + c.Param("module"))
}

func SetSettingsPerModulePerKey(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI)
}
