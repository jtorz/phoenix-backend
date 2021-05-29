package baseservice

import "context"

// Agent has the information of the user that is executing an operation,
// and an AuthService to retrive their privileges.
type Agent struct {
	UserID string
	AgentInfoService
	AuthService
}

// NewAgent creates a new agent with the services.
func NewAgent(userId string, infoSvc AgentInfoService, authSvc AuthService) *Agent {
	return &Agent{
		UserID:           userId,
		AgentInfoService: infoSvc,
		AuthService:      authSvc,
	}
}

// AuthService is used to retrieve the privileges of the agent.
type AuthService interface {
	GetPrivilegeByPriority(...string) (string, error)
	HasPrivilege(string) (bool, error)
	IsAdmin() (bool, error)
}

// AgentInfoService is used to retrieve the information of the agent.
type AgentInfoService interface {
	GetInfo(context.Context) (AgentInfo, error)
}

//  AgentInfo holds the general information of the agent.
type AgentInfo struct {
	ID         string
	Name       string
	MiddleName string
	LastName   string
	Email      string
	Username   string
}

type anonym struct{}

// Creates a new anonym agent.
func NewAgentAnonym() *Agent {
	a := anonym{}
	return &Agent{
		AgentInfoService: a,
		AuthService:      a,
	}
}

func (anonym) GetInfo(context.Context) (AgentInfo, error) {
	return AgentInfo{}, nil
}
func (anonym) GetPrivilegeByPriority(...string) (string, error) {
	return "", nil
}

func (anonym) HasPrivilege(string) (bool, error) {
	return false, nil
}

func (anonym) IsAdmin() (bool, error) {
	return false, nil
}
