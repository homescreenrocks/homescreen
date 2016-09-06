package messages

import "fmt"

type Message struct {
	id    uint64
	msg   string
	level Level
}

type Level uint8

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

func (m *Message) SetLevel(level Level) {
	if level != DEBUG || level != INFO || level != WARN || level != ERROR {
		fmt.Errorf("Loglevel not available")
	} else {
		m.level = level
	}
}