package unit

import (
	"github.com/bernardolins/clustereasy/setup/types"
	"strings"
)

type Unit struct {
	name    string
	command string
	content string
	runtime bool
	enable  bool
}

func New(name string) *Unit {
	u := new(Unit)

	u.name = name
	//u.command = unit.UnitCommand()
	//u.runtime = unit.UnitRuntime()
	//u.enable = unit.UnitEnable()

	return u
}

func (u Unit) GetName() string {
	return u.name
}

func (u Unit) GetCommand() string {
	return u.command
}

func (u Unit) GetContent() string {
	return u.content
}

func (u Unit) GetRuntime() bool {
	return u.runtime
}

// ---------------------

func (u *Unit) AddCommand(command string) {
	u.command = command
}

func (u *Unit) Enable() {
	u.enable = true
}

func (u *Unit) OnRuntime() {
	u.runtime = true
}

func (u *Unit) SetContent(content string) {
	u.content = content
}

func (u *Unit) Configure(unit types.Unit) {
	u.command = unit.UnitCommand()
	u.runtime = unit.UnitRuntime()
	u.enable = unit.UnitEnable()
}

// ---------------------

func (u Unit) ContentLines() []string {
	lines := strings.Split(u.content, "\n")
	return lines
}
