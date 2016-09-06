package types

type Module struct {
	ID          uint64
	Name        string
	Version     string
	Description string
	Dir         string
	Settings    map[string]Setting
}

type Setting struct {
	Default        string
	Type           string
	Mandatory      bool
	Description    string
	PossibleValues []string
	Value          interface{}
}
