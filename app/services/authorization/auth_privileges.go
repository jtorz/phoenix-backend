package authorization

type Privileges struct {
	Roles      []string
	Privileges []Privilege
}

type Privilege struct {
	Method string
	Route  string
	Key    string
}

func (p Privilege) GetKey() string {
	return p.Key
}
