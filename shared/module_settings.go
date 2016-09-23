package shared

type ModuleSettings []ModuleSetting

type ModuleSetting struct {
	Name           string      `json:"name"`
	Default        interface{} `json:"default"`
	Type           string      `json:"type"`
	Mandatory      bool        `json:"mandatory"`
	Description    string      `json:"description"`
	PossibleValues []string    `json:"possible-values,omitempty"`
	Value          interface{} `json:"value"`
}
