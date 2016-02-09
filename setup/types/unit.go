package types

type Unit struct {
	Name    string `yaml: "name"`
	Command string `yaml: "command"`
}

func (u *Unit) UnitName() string {
	return u.Name
}

func (u *Unit) UnitCommand() string {
	return u.Command
}
