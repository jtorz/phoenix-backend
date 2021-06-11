package coremodel

// Privileges slice.
type Privileges []Privilege

// Privilege action that an user is allowed to do.
type Privilege struct {
	Role   Role
	Action Action
}

func (p Privilege) Code() string {
	return p.Action.Key()
}

// GetCodes returns a slice with all the codes of the privileges
func (p Privileges) Codes() []string {
	s := make([]string, len(p))
	for i := range p {
		s[i] = p[i].Code()
	}
	return s
}
