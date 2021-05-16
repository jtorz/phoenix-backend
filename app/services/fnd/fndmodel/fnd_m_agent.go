package fndmodel

// AgentGetter interface to retrieve the information of an agent.
type AgentGetter interface {
	// Returns the information of an agent.
	GetAgent() (Agent, error)
}

// Agent represents a user in a request context.
type Agent interface {
	// UserID returns the id of the user.
	UserID() string
	// Privileges returns the privileges that the user has.
	Privileges() (Privileges, error)
	// ToUser returns the user Struct
	ToUser() User
}

// Anonymous anonymous user.
var Anonymous = anonymous{}

type anonymous struct{}

// UserID returns the id of the user. An anonymous user does't have an ID.
func (anonymous) UserID() string {
	return ""
}

// Privileges returns the privileges that the user has. An anonyous user does't have any privileges.
func (anonymous) Privileges() (Privileges, error) {
	return []Privilege{}, nil
}

// ToUser returns the user Struct (empty).
func (anonymous) ToUser() User {
	return User{Roles: []Role{}}
}
