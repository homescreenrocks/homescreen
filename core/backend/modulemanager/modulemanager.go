package modulemanager

import (
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/homescreen/core/backend/storage"
	"github.com/homescreenrocks/homescreen/shared"
)

const MODULES_FOLDER = "./modules"

type ModuleManager struct {
	modules  map[string]*shared.Module
	storage  *storage.Storage
	settings ModuleManagerSettings
}

type ModuleManagerSettings struct {
	execMode bool
}

// New creates ModuleManager instance
func New(storage *storage.Storage, execmode bool) ModuleManager {
	mm := ModuleManager{}
	mm.modules = make(map[string]*shared.Module)

	mm.storage = storage

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
			c.JSON(400, shared.HttpError{"Decoding HTTP body failed.", err})
			return
		}

		if strings.TrimSpace(req.ModuleURL) == "" {
			c.JSON(400, shared.HttpError{"Attribute 'module-url' is missing.", nil})
			return
		}

		if mm.GetModule(req.Metadata.Name) != nil {
			c.JSON(423, shared.HttpError{"Module already registered.", nil})
			return
		}

		mm.AddModule(&req)
	})
}

// GetAllModules returns all Modules as map[string]types.Module
func (mm *ModuleManager) GetAllModules() []*shared.Module {
	modules := make([]*shared.Module, 0)
	for key, _ := range mm.modules {
		modules = append(modules, mm.GetModule(key))
	}

	return modules
}

// GetModule returns specified module or an empty one if module does not exist
func (mm *ModuleManager) GetModule(id string) *shared.Module {
	val, ok := mm.modules[id]
	if !ok {
		return nil
	}

	cpy := *val

	for i, setting := range cpy.Settings {
		key := path.Join("module", id, setting.Name)
		err := mm.storage.Get(key, &setting.Value)

		if err != nil {
			if _, ok := err.(storage.KeyNotFound); !ok {
				log.Printf("Error getting data from storage: %v", err)
			}
			setting.Value = setting.Default
		}
		cpy.Settings[i] = setting
	}

	return &cpy
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

func (mm *ModuleManager) AddModule(m *shared.Module) {
	mm.modules[m.Metadata.ID] = m
}
