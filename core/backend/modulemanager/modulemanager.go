package modulemanager

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

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
	group.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, (*mm).GetAllModules())
	})

	group.POST("/", func(c *gin.Context) {
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

		if mm.GetModule(req.Metadata.ID) != nil {
			c.JSON(423, shared.HttpError{"Module already registered.", nil})
			return
		}

		mm.AddModule(&req)
	})

	group.Any("/:module/proxy/*path", func(c *gin.Context) {
		module := mm.GetModule(c.Param("module"))
		if module == nil {
			c.JSON(404, shared.HttpError{"Module not found.", nil})
			return
		}

		target, err := url.Parse(module.ModuleURL)
		if module == nil {
			c.JSON(500, shared.HttpError{"Error connecting to module.", err})
			return
		}

		oldPrefix := fmt.Sprintf("%s/%s/proxy", group.BasePath(), module.Metadata.ID)
		newPrefix := "/module"

		var handler http.Handler
		handler = httputil.NewSingleHostReverseProxy(target)
		handler = HTTPAddPrefix(newPrefix, handler)
		handler = http.StripPrefix(oldPrefix, handler)

		wrapper := gin.WrapH(handler)
		wrapper(c)
	})
}

// GetAllModules returns all Modules as map[string]types.Module
func (mm *ModuleManager) GetAllModules() []*shared.Module {
	modules := make([]*shared.Module, 0)
	for key := range mm.modules {
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
	for _ = range files {
		//mm.readModuleConfig(f)
	}

}

func (mm *ModuleManager) Count() int {
	return len(mm.modules)
}

func (mm *ModuleManager) AddModule(m *shared.Module) {
	mm.modules[m.Metadata.ID] = m
	mm.AddWatchdog(m.Metadata.ID)

}

func (mm *ModuleManager) AddWatchdog(id string) {
	go func() {
		flushTicker := time.NewTicker(5 * time.Second)
		for range flushTicker.C {
			log.Print("Watchdog of Module " + id)
			log.Print(mm.GetModule(id).ModuleURL)
			resp, err := http.Get(mm.GetModule(id).ModuleURL + "/hello")
			if err != nil || resp.StatusCode != 200 {
				log.Print("Module is not available, I need to remove it")
			}
		}
	}()
}
