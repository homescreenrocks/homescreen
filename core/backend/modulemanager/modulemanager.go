package modulemanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/homescreen/core/backend/types"
)

const MODULES_FOLDER = "./modules"

type ModuleManager struct {
	modules  map[string]types.Module
	settings ModuleManagerSettings
}

type ModuleManagerSettings struct {
	execMode bool
}

// New creates ModuleManager instance
func New(execmode bool) ModuleManager {
	mm := ModuleManager{}
	mm.modules = make(map[string]types.Module)

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
		var m types.Module
		dec := json.NewDecoder(c.Request.Body)
		dec.Decode(&m)
		mm.AddModule(m)
	})
}

// GetAllModules returns all Modules as map[string]types.Module
func (mm *ModuleManager) GetAllModules() map[string]types.Module {
	return mm.modules
}

// GetModule returns specified module or an empty one if module does not exist
func (mm *ModuleManager) GetModule(key string) types.Module {
	if val, ok := mm.modules[key]; ok {
		return val
	}

	return types.Module{}
}

// ScanForModules scans for modules in modules folder
func (mm *ModuleManager) ScanForModules() {
	folder, _ := filepath.Abs(MODULES_FOLDER)
	// search for module.json files in the modules folder
	files, _ := filepath.Glob(folder + "/**/module.json")
	for _, f := range files {
		mm.readModuleConfig(f)
	}

}

func (mm *ModuleManager) Count() int {
	return len(mm.modules)
}

func (mm *ModuleManager) readModuleConfig(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("Couldn't read file: %s\n", file)
	}

	var m types.Module
	err = json.Unmarshal(content, &m)

	m.Dir = filepath.Dir(file)

	mm.modules[m.Name] = m
}

func (mm *ModuleManager) AddModule(m types.Module) {
	mm.modules[m.Name] = m
}
