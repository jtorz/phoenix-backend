package authorization

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {
	svcs := []JWTSvc{
		[]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"),
		[]byte("123456789"),
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]byte{255, 255, 255},
		[]byte{0},
	}
	for _, svc := range svcs {
		original := &baseservice.JWTData{
			ID: "591e58a6-e04e-5239-a0e9-7ee3c4ca2423",
		}
		s, err := svc.NewJWT(*original)
		assert.Nil(t, err)
		if err != nil {
			return
		}

		decoded, err := svc.authJWT(s)
		assert.Nil(t, err)
		if err != nil {
			return
		}
		assert.Equal(t, original, decoded)
	}
}

func TestExpiredToken(t *testing.T) {
	svc := JWTSvc([]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"))
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjI0MzAzMDIsInVzZXJJRCI6IjU5MWU1OGE2LWUwNGUtNTIzOS1hMGU5LTdlZTNjNGNhMjQyMyJ9.eHoVa9dD_tIvsyJ_yHG-4JJGX73bqkLanR17OliI5Ho"
	decoded, err := svc.authJWT(s)
	assert.Nil(t, decoded)
	assert.NotNil(t, err)
}

func TestParseBadJWT(t *testing.T) {
	svc := JWTSvc([]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"))
	original := &baseservice.JWTData{
		ID: "591e58a6-e04e-5239-a0e9-7ee3c4ca2423",
	}
	s, err := svc.NewJWT(*original)
	assert.Nil(t, err)
	if err != nil {
		return
	}

	decoded, err := svc.authJWT(s + "a")
	assert.Nil(t, decoded)
	assert.NotNil(t, err)

	decoded, err = svc.authJWT("")
	assert.Nil(t, decoded)
	assert.NotNil(t, err)
}

func TestPanicOnEmpty(t *testing.T) {
	assert.Panics(t, func() {
		svc := JWTSvc([]byte{})
		svc.NewJWT(baseservice.JWTData{})
	})
}
func TestBearerToken(t *testing.T) {
	svc := JWTSvc([]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"))
	original := &baseservice.JWTData{
		ID: "591e58a6-e04e-5239-a0e9-7ee3c4ca2423",
	}
	s, err := svc.NewJWT(*original)
	assert.Nil(t, err)
	if err != nil {
		return
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		decoded, err := svc.AuthJWT(r)
		assert.Nil(t, err)
		if err != nil {
			return
		}
		assert.Equal(t, original, decoded)
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+s)
	rec := httptest.NewRecorder()

	handler(rec, req)
}

func TestNotBearerToken(t *testing.T) {
	svc := JWTSvc([]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"))
	original := &baseservice.JWTData{
		ID: "591e58a6-e04e-5239-a0e9-7ee3c4ca2423",
	}
	s, err := svc.NewJWT(*original)
	assert.Nil(t, err)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
	}
	req.Header.Set("Authorization", s)

	decoded, err := svc.AuthJWT(req)

	assert.Nil(t, decoded)
	assert.NotNil(t, err)
}

func TestEmptyToken(t *testing.T) {
	svc := JWTSvc([]byte(":bu}V?8UAbc/x,rZ;+pTpZB:R+HEX(9&rTXj8?2h:9UU/;a;{3p,QB6?E&MQ"))

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create a request %v", err)
	}

	decoded, err := svc.AuthJWT(req)
	assert.Nil(t, decoded)
	assert.NotNil(t, err)
}
