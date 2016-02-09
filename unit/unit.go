package unit

type Unit struct {
	name    string
	command string
	content string
}

func New(name, command string) *Unit {
	u := new(Unit)
	u.name = name
	u.command = command

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
