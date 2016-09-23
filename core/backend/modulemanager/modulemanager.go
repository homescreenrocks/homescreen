package modulemanager

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/homescreen/shared"
)

const MODULES_FOLDER = "./modules"

type ModuleManager struct {
	modules  map[string]shared.Module
	settings ModuleManagerSettings
}

type ModuleManagerSettings struct {
	execMode bool
}

// New creates ModuleManager instance
func New(execmode bool) ModuleManager {
	mm := ModuleManager{}
	mm.modules = make(map[string]shared.Module)

	// set settings
	mm.settings.execMode = execmode
	log.Println(mm)
	return mm
}

func (mm *ModuleManager) RegisterRouterGroup(group *gin.RouterGroup) {
	group.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, (*mm).GetAllModules())
	})

	group.POST("/register", func(c *gin.Context) {
		var req shared.Module
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(400, HttpError{"Decoding HTTP body failed.", err})
			return
		}

		if strings.TrimSpace(req.ModuleURL) == "" {
			c.JSON(400, HttpError{"Attribute 'module-url' is missing.", nil})
			return
		}

		if mm.GetModule(req.Metadata.Name) != nil {
			c.JSON(423, HttpError{"Module already registered.", nil})
			return
		}

		mm.AddModule(req)
	})
}

// GetAllModules returns all Modules as map[string]types.Module
func (mm *ModuleManager) GetAllModules() map[string]shared.Module {
	return mm.modules
}

// GetModule returns specified module or an empty one if module does not exist
func (mm *ModuleManager) GetModule(key string) *shared.Module {
	if val, ok := mm.modules[key]; ok {
		return &val
	}

	return nil
}

// ScanForModules scans for modules in modules folder
func (mm *ModuleManager) ScanForModules() {
	folder, _ := filepath.Abs(MODULES_FOLDER)
	// search for module.json files in the modules folder
	files, _ := filepath.Glob(folder + "/**/module.json")
	for _, _ = range files {
		//mm.readModuleConfig(f)
	}

}

func (mm *ModuleManager) Count() int {
	return len(mm.modules)
}

func (mm *ModuleManager) AddModule(m shared.Module) {
	mm.modules[m.Metadata.ID] = m
}
