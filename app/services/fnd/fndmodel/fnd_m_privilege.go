package fndmodel

// Privileges slice.
type Privileges []Privilege

// Privilege action that an user is allowed to do.
type Privilege struct {
	Role   Role   `json:"role"`
	Action Action `json:"action"`
}

func (p Privilege) Code() string {
	return p.Action.Module.ID + "." + p.Action.ID
}

// GetCodes returns a slice with all the codes of the privileges
func (p Privileges) Codes() []string {
	s := make([]string, len(p))
	for i := range p {
		s[i] = p[i].Code()
	}
	return s
}
