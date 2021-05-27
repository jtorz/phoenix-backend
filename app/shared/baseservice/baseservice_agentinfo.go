package baseservice

import "context"

type AgentInfo struct {
	ID         string
	Name       string
	MiddleName string
	LastName   string
	Email      string
	Username   string
}

type Anonym struct{}

func (Anonym) GetInfo(ctx context.Context) (AgentInfo, error) {
	return AgentInfo{}, nil
}
