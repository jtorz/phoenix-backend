package authorization_test

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/config/testconfig"
	"github.com/jtorz/phoenix-backend/app/services/authorization"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
	"github.com/stretchr/testify/assert"

	"github.com/jtorz/phoenix-backend/app/shared/lex"
)

var testAuthSvc = struct {
	dbTx *sql.Tx

	jwtSvc authorization.JWTSvc

	adminUser  testUserData
	normalUser testUserData

	config *testconfig.Config
	ctx    context.Context
}{}

type testUserData struct {
	id  string
	jwt string
}

func TestMain(m *testing.M) {
	var err error
	testAuthSvc.config, err = testconfig.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	testAuthSvc.dbTx, err = testAuthSvc.config.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	conn := testAuthSvc.config.Redis.Get()
	conn.Do("FLUSHALL")
	conn.Close()

	err = prepareAuthorizationTest()
	if err != nil {
		log.Fatal(err)
	}
	m.Run()

	testAuthSvc.dbTx.Rollback()
}

func prepareAuthorizationTest() (err error) {
	testAuthSvc.ctx = ctxinfo.SetLoggingLevelC(context.Background(), testAuthSvc.config.LoggingLevel)
	testAuthSvc.jwtSvc = authorization.JWTSvc([]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"))

	if err = auxTest.insertModule("_TST_MOD_"); err != nil {
		return
	}
	if err = auxTest.insertActions("_TST_MOD_", "A", "B", "C", "D"); err != nil {
		return
	}
	if err = auxTest.insertActionRoute("_TST_MOD_", "A", "GET", "/api/test/get"); err != nil {
		return
	}
	if err = auxTest.insertRole("_TESTROLE_"); err != nil {
		return
	}

	if err = auxTest.insertPrivileges("_TESTROLE_", "_TST_MOD_.A", "_TST_MOD_.B", "_TST_MOD_.C"); err != nil {
		return
	}

	if err = prepareNormalUser(); err != nil {
		return
	}
	if err = prepareAdmin(); err != nil {
		return
	}
	return nil
}

type auxTestStruct struct{}

var auxTest = auxTestStruct{}

func (auxTestStruct) insertUser(userID string) (err error) {
	_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CoreUser).Rows(goqu.Record{
		lex.CoreUser.UseID:         userID,
		lex.CoreUser.UseName:       userID,
		lex.CoreUser.UseMiddleName: userID,
		lex.CoreUser.UseLastName:   userID,
		lex.CoreUser.UseEmail:      userID,
		lex.CoreUser.UseUsername:   userID,
		lex.CoreUser.UseStatus:     2,
	}))
	return
}

func (auxTestStruct) insertModule(moduleID string) (err error) {
	_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CoreModule).Rows(goqu.Record{
		lex.CoreModule.ModID:          moduleID,
		lex.CoreModule.ModName:        moduleID,
		lex.CoreModule.ModDescription: moduleID,
		lex.CoreModule.ModOrder:       1,
		lex.CoreModule.ModStatus:      2,
	}))
	return
}

func (auxTestStruct) insertActions(moduleID string, actions ...string) (err error) {
	for _, a := range actions {
		_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CoreAction).Rows(goqu.Record{
			lex.CoreAction.ActModuleID:    moduleID,
			lex.CoreAction.ActActionID:    a,
			lex.CoreAction.ActName:        a,
			lex.CoreAction.ActDescription: a,
			lex.CoreAction.ActOrder:       1,
			lex.CoreAction.ActStatus:      2,
		}))
		if err != nil {
			return
		}
	}
	return
}

func (auxTestStruct) insertActionRoute(moduleID, actionID, method, route string) (err error) {
	_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CoreActionRoute).Rows(goqu.Record{
		lex.CoreActionRoute.AcrModuleID: moduleID,
		lex.CoreActionRoute.AcrActionID: actionID,
		lex.CoreActionRoute.AcrMethod:   method,
		lex.CoreActionRoute.AcrRoute:    route,
	}))
	return
}

func (auxTestStruct) insertRole(rolID string) (err error) {
	_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CoreRole).Rows(goqu.Record{
		lex.CoreRole.RolID:          rolID,
		lex.CoreRole.RolName:        rolID,
		lex.CoreRole.RolDescription: rolID,
		lex.CoreRole.RolIcon:        rolID,
		lex.CoreRole.RolStatus:      2,
	}))
	return
}

func (auxTestStruct) insertPrivileges(rolID string, privs ...string) (err error) {
	for _, priv := range privs {
		act := strings.Split(priv, ".")
		_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CorePrivilege).Rows(goqu.Record{
			lex.CorePrivilege.PriRoleID:   "_TESTROLE_",
			lex.CorePrivilege.PriModuleID: act[0],
			lex.CorePrivilege.PriActionID: act[1],
		}))
		if err != nil {
			return err
		}
	}
	return nil
}

func (auxTestStruct) insertUserRole(userId, rolID string) (err error) {
	_, err = lex.DoInsert(testAuthSvc.ctx, testAuthSvc.dbTx, lex.NewInsert(lex.T.CoreUserRole).Rows(goqu.Record{
		lex.CoreUserRole.UsrUserID: userId,
		lex.CoreUserRole.UsrRoleID: rolID,
	}))
	return
}

// ************************************************************************************
// ************************************************************************************
// ************************************ ADMIN USER ************************************
// ************************************************************************************
// ************************************************************************************

func prepareAdmin() (err error) {
	userID := "591ea8a6-e049-5279-a049-75e3c4ca2423"
	testAuthSvc.adminUser.id = userID
	testAuthSvc.adminUser.jwt, err = testAuthSvc.jwtSvc.NewJWT(baseservice.JWTData{ID: userID})
	if err != nil {
		return
	}

	if err = auxTest.insertUser(userID); err != nil {
		return
	}
	if err = auxTest.insertUserRole(userID, baseservice.RoleAdmin); err != nil {
		return
	}
	return nil
}

func TestAuthSvcAdmin(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/test/get", nil)
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+testAuthSvc.adminUser.jwt)

	auth, err := authorization.NewAuthService(req, testAuthSvc.jwtSvc, testAuthSvc.dbTx, testAuthSvc.config.Redis)
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
		return
	}
	err = auth.CheckAuthorization(testAuthSvc.ctx, "GET", "/api/test/get")
	assert.Nil(t, err)

	// First call from DB
	ok, err := auth.IsAdmin(testAuthSvc.ctx)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.True(t, ok)

	// Second call from cache
	ok, err = auth.IsAdmin(testAuthSvc.ctx)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.True(t, ok)
}

// ************************************************************************************
// ************************************************************************************
// ************************************ NORMAL USER ***********************************
// ************************************************************************************
// ************************************************************************************

func prepareNormalUser() (err error) {
	userID := "77e8a8e5-2d69-4c74-974d-25381f1ffd2d"
	testAuthSvc.normalUser.id = userID
	testAuthSvc.normalUser.jwt, err = testAuthSvc.jwtSvc.NewJWT(baseservice.JWTData{ID: userID})
	if err != nil {
		return
	}
	if err = auxTest.insertUser(userID); err != nil {
		return
	}

	if err = auxTest.insertUserRole(userID, "_TESTROLE_"); err != nil {
		return
	}

	return nil
}
func TestAuthSvcNormalUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/test/get", nil)
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+testAuthSvc.normalUser.jwt)

	auth, err := authorization.NewAuthService(req, testAuthSvc.jwtSvc, testAuthSvc.dbTx, testAuthSvc.config.Redis)
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
		return
	}
	err = auth.CheckAuthorization(testAuthSvc.ctx, "GET", "/api/test/get")
	assert.Nil(t, err)

	// First call from DB
	ok, err := auth.IsAdmin(testAuthSvc.ctx)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.False(t, ok)

	// Second call from cache
	ok, err = auth.IsAdmin(testAuthSvc.ctx)
	assert.Nil(t, err)
	if err != nil {
		return
	}
	assert.False(t, ok)
}

func TestAuthSvcNormalUserForbbiden(t *testing.T) {
	req, err := http.NewRequest("POST", "/api/test/something", nil)
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+testAuthSvc.normalUser.jwt)

	auth, err := authorization.NewAuthService(req, testAuthSvc.jwtSvc, testAuthSvc.dbTx, testAuthSvc.config.Redis)
	assert.Nil(t, err)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
		return
	}
	err = auth.CheckAuthorization(testAuthSvc.ctx, "POST", "/api/test/something")
	assert.True(t, baseerrors.IsErrPrivilege(err))

	// First call from DB
	ok, err := auth.IsAdmin(testAuthSvc.ctx)
	assert.Nil(t, err)
	assert.False(t, ok)

	// Second call from cache
	ok, err = auth.IsAdmin(testAuthSvc.ctx)
	assert.Nil(t, err)
	assert.False(t, ok)

	ok, err = auth.HasPrivilege(testAuthSvc.ctx, "_TST_MOD_.A")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = auth.HasPrivilege(testAuthSvc.ctx, "_TST_MOD_.THIS_SHOULD_BE_FALSE")
	assert.Nil(t, err)
	assert.False(t, ok)

	priv, err := auth.GetPrivilegeByPriority(testAuthSvc.ctx, "_TST_MOD_.B", "_TST_MOD_.A")
	assert.Nil(t, err)
	assert.Equal(t, "_TST_MOD_.B", priv)

	priv, err = auth.GetPrivilegeByPriority(testAuthSvc.ctx, "_TST_MOD_.D", "_TST_MOD_.A")
	assert.Nil(t, err)
	assert.Equal(t, "_TST_MOD_.A", priv)

	priv, err = auth.GetPrivilegeByPriority(testAuthSvc.ctx, "_TST_MOD_.NOPE1", "_TST_MOD_.NOPE2")
	assert.Nil(t, err)
	assert.Equal(t, "", priv)
}
