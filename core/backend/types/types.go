package types

type Setting struct {
	Default        string
	Type           string
	Mandatory      bool
	Description    string
	PossibleValues []string
	Value          interface{}
}
