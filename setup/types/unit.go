package types

type Unit struct {
	Name    string `yaml: "name"`
	Command string `yaml: "command,omitempty"`
	Runtime bool   `yaml: "runtime,omitempty"`
	Enable  bool   `yaml: "enable,omitempty"`
}

func (u *Unit) UnitName() string {
	return u.Name
}

func (u *Unit) UnitCommand() string {
	return u.Command
}

func (u *Unit) UnitRuntime() bool {
	return u.Runtime
}

func (u *Unit) UnitEnable() bool {
	return u.Enable
}
