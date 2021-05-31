package baseservice

import "context"

// RoleAdmin system admin role ID.
const RoleAdmin string = "SYS_ADM"

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
	// GetPrivilegeByPriority returns the privilege with the highest priority.
	// The priority is given by the order in which the parameters are passed to the function.
	// If the agent doesn't have any of the privileges an empty string and a nil err is returned.
	//
	//	 GetPrivilegeByPriority("FIRST_PARAMETER_HAS_THE_HIGHEST", "OTHER", "ANOTHER", "LAST_PARAMETER_HAS_THE_LOWEST")
	//
	// For example lets consider an agent with the privileges:
	// ["QUERY_ALL", "QUERY_ACTIVE", "QUERY_SPECIFIC", "OTHER1", "OTHER2"]
	//
	//	s, _ := GetPrivilegeByPriority("QUERY_ALL","QUERY_ACTIVE","QUERY_SPECIFIC")
	//	// s == "QUERY_ALL"
	//	s, _ := GetPrivilegeByPriority("OTHER3","OTHER2","OTHER1")
	//	// s == "OTHER2"
	//	s, _ := GetPrivilegeByPriority("FOO3","FOO2","FOO1")
	//	// s == ""
	GetPrivilegeByPriority(context.Context, ...string) (string, error)
	// HasPrivilege Checks if the agent has an specific privilege.
	HasPrivilege(context.Context, string) (bool, error)
	// IsAdmin Checks if the agent is the general admin.
	IsAdmin(context.Context) (bool, error)
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
func (anonym) GetPrivilegeByPriority(context.Context, ...string) (string, error) {
	return "", nil
}

func (anonym) HasPrivilege(context.Context, string) (bool, error) {
	return false, nil
}

func (anonym) IsAdmin(context.Context) (bool, error) {
	return false, nil
}
