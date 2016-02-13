package unit

import (
	"github.com/bernardolins/clustereasy/setup/types"
	"strings"
)

type Unit struct {
	name    string
	params  map[string]interface{}
	command string
	runtime bool
	enable  bool
	content string
	dropins []DropIn
}

func New(name string) *Unit {
	u := new(Unit)

	u.name = name

	u.params = make(map[string]interface{})
	//u.command = unit.UnitCommand()
	//u.runtime = unit.UnitRuntime()
	//u.enable = unit.UnitEnable()

	return u
}

func (u Unit) Name() string {
	return u.name
}

func (u Unit) Command() string {
	return u.command
}

func (u Unit) Content() string {
	return u.content
}

func (u Unit) Runtime() bool {
	return u.runtime
}

func (u Unit) Enable() bool {
	return u.enable
}

func (u Unit) Parameters() map[string]interface{} {
	return u.params
}

// ---------------------

func (u *Unit) AddCommand(command string) {
	u.command = command
}

func (u *Unit) SetEnable() {
	u.enable = true
}

func (u *Unit) OnRuntime() {
	u.runtime = true
}

func (u *Unit) SetContent(content string) {
	u.content = content
}

func (u *Unit) Configure(unit types.Unit) {
	u.params["command"] = unit.UnitCommand()
	u.params["enable"] = unit.UnitEnable()
	u.params["runtime"] = unit.UnitRuntime()
}

// ---------------------

func (u Unit) ContentLines() []string {
	lines := strings.Split(u.content, "\n")
	return lines
}
