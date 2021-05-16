package ctxinfo

import "context"

func PrintLog(ctx context.Context) bool {
	debug, _ := ctx.Value("DEBUG").(bool)
	qa, _ := ctx.Value("QA").(bool)
	return debug || qa
}

type Agent struct {
	UserID string
	AuthService
}

type agentContextKey struct {
}

type AuthService interface {
	GetPrivileges() (Privilege, error)
}

type Privilege struct {
	Method string
	Route  string
	Key    string
}

func SetAgent(c context.Context, userID string, authSvc AuthService) context.Context {
	return context.WithValue(c, agentContextKey{}, Agent{
		UserID:      userID,
		AuthService: authSvc,
	})
}
