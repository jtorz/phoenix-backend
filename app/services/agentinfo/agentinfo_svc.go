package agentinfo

import (
	"context"
	"database/sql"
	"sync"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
	"github.com/jtorz/phoenix-backend/app/shared/lex"
)

// Service authorization module service.
//
// Implements ctxinfo.AgentInfoService interface to retrieve the user data.
type Service struct {
	db            *sql.DB
	mux           *sync.Mutex
	userRetrieved bool
	userID        string

	agentInfo baseservice.AgentInfo
}

func NewService(db *sql.DB, userID string) *Service {
	return &Service{db: db, userID: userID, mux: &sync.Mutex{}}
}

func (svc *Service) GetInfo(ctx context.Context) (baseservice.AgentInfo, error) {
	if err := svc.getUser(ctx); err != nil {
		return baseservice.AgentInfo{}, err
	}
	return svc.agentInfo, nil
}

// getUser searchs the user with the given filters.
func (svc *Service) getUser(ctx context.Context) error {
	svc.mux.Lock()
	defer svc.mux.Unlock()
	if svc.userRetrieved {
		return nil
	}
	query := lex.NewSelect(
		lex.CoreUser.UseEmail,
		lex.CoreUser.UseUsername,
		lex.CoreUser.UseName,
		lex.CoreUser.UseMiddleName,
		lex.CoreUser.UseLastName,
	).
		From(lex.T.CoreUser).
		Where(goqu.C(lex.CoreUser.UseID).Eq(svc.userID))

	row, err := lex.QueryRowContext(ctx, svc.db, query)
	if err != nil {
		lex.DebugErr(ctx, err)
		return err
	}
	err = row.Scan(
		&svc.agentInfo.Email,
		&svc.agentInfo.Username,
		&svc.agentInfo.Name,
		&svc.agentInfo.MiddleName,
		&svc.agentInfo.LastName,
	)
	if err != nil {
		lex.DebugErr(ctx, err)
		return err
	}
	svc.agentInfo.ID = svc.userID
	svc.userRetrieved = true
	return nil
}
