package shared

type Module struct {
	ModuleURL string         `json:"module-url"`
	Metadata  ModuleMetadata `json:"metadata"`
	Settings  ModuleSettings `json:"settings"`
}
