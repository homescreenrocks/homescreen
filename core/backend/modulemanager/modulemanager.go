package modulemanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/homescreenrocks/homescreen/core/backend/types"
)

const MODULES_FOLDER = "./modules"

type ModuleManager struct {
	modules map[string]types.Module
}

// New creates ModuleManager instance
func New() ModuleManager {
	mm := ModuleManager{}
	mm.modules = make(map[string]types.Module)

	return mm
}

// GetAllModules returns all Modules as map[string]types.Module
func (mm *ModuleManager) GetAllModules() map[string]types.Module {
	return mm.modules
}

//GetAllModulesAsJSON return JSON
func (mm *ModuleManager) GetAllModulesAsJSON() []byte {
	ret, _ := json.Marshal(mm.modules)
	return ret
}

//GetModuleAsJSON return JSON for specified module
func (mm *ModuleManager) GetModuleAsJSON(key string) []byte {
	ret, _ := json.Marshal(mm.modules[key])
	return ret
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
	//fmt.Println(files)
	for _, f := range files {
		//fmt.Println("looping the files")
		//fmt.Println(f)
		mm.readModuleConfig(f)
	}

}

func (mm *ModuleManager) moduleExists(m string) bool {
	return false
}

func (mm *ModuleManager) Count() int {
	return len(mm.modules)
}

func (mm *ModuleManager) readModuleConfig(file string) {
	//fmt.Println("trying to read file", file)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("Couldn't read file: %s\n", file)
	}

	var m types.Module
	err = json.Unmarshal(content, &m)
	//fmt.Println(m)

	//x := m.(map[string]interface{})
	// x["apikey"]
	/*
		fmt.Println(m.Settings)
		fmt.Println(m.Settings["apikey"])
		fmt.Println(m.Settings["apikey"].Default)
		fmt.Println(m.Settings["apikey"].Description)
		fmt.Println(m.Settings["apikey"].Mandatory)
		fmt.Println(m.Settings["apikey"].PossibleValues)
		fmt.Println(m.Settings["apikey"].Type)

		fmt.Println(m.Settings["language"])
		fmt.Println(m.Settings["language"].Default)
		fmt.Println(m.Settings["language"].Description)
		fmt.Println(m.Settings["language"].Mandatory)
		fmt.Println(m.Settings["language"].PossibleValues)
		fmt.Println(m.Settings["language"].Type)
	*/
	m.Dir = filepath.Dir(file)
	//fmt.Printf("%T %T", x.(map[string]interface{}))
	//log.Println("done with reading module config", file)
	mm.modules[m.Name] = m
}

func (mm *ModuleManager) AddModule(m types.Module) {
	mm.modules[m.Name] = m
}
