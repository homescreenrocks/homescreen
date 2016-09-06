package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/homescreenrocks/homescreen/core/backend/modulemanager"
	"github.com/homescreenrocks/homescreen/core/backend/types"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	ModuleManager *modulemanager.ModuleManager
}

func New(mm *modulemanager.ModuleManager) *Routes {
	return &Routes{
		ModuleManager: mm,
	}
}

func (r *Routes) GetModules(c *gin.Context) {
	mm := r.ModuleManager
	c.JSON(http.StatusOK, (*mm).GetAllModules())
}

func (r *Routes) GetModulesAsJSON(c *gin.Context) {
	mm := r.ModuleManager
	c.JSON(http.StatusOK, (*mm).GetAllModules())
}

func (r *Routes) FindModules(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI)
}

func (r *Routes) ExecuteModules(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI + "var is: " + c.Param("module"))
}

func (r *Routes) GetSettings(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI + " simple types")
}

func (r *Routes) SetSettingsPerModule(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI + "var is: " + c.Param("module"))
}

func (r *Routes) SetSettingsPerModulePerKey(c *gin.Context) {
	log.Println("the request path is: " + c.Request.RequestURI)
}

func (r *Routes) RegisterModule(c *gin.Context) {
	mm := r.ModuleManager
	var m types.Module
	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&m)
	mm.AddModule(m)
}
