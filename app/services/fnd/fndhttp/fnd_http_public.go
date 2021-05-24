package fndhttp

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndbiz"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// httpPublic http handler component.
type httpPublic struct {
	DB      *sql.DB
	JwtSvc  baseservice.JWTGeneratorSvc
	MailSvc baseservice.MailSenderSvc
}

func newHttpPublic(db *sql.DB, jwtSvc baseservice.JWTGeneratorSvc, mailSvc baseservice.MailSenderSvc) httpPublic {
	return httpPublic{
		DB:      db,
		JwtSvc:  jwtSvc,
		MailSvc: mailSvc,
	}
}

// Signup http handler.
func (handler httpPublic) Signup() httphandler.HandlerFunc {
	type Req struct {
		Name       string `binding:"required"`
		MiddleName string `binding:"required"`
		LastName   string `binding:"required"`
		Email      string `binding:"required,email"`
		Username   string `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		u := fndmodel.User{
			Name:       req.Name,
			MiddleName: req.MiddleName,
			LastName:   req.LastName,
			Email:      req.Email,
			Username:   req.Username,
		}

		tx := c.BeginTx(handler.DB)
		biz := fndbiz.NewBizUser()
		err := biz.New(c, tx.Tx, handler.MailSvc, &u)
		if baseerrors.IsErrAuth(err) {
			c.Status(http.StatusUnauthorized)
			tx.Rollback(c)
			return
		} else if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}

// Login http handler.
func (handler httpPublic) Login() httphandler.HandlerFunc {
	type Req struct {
		User     string `binding:"required"`
		Password string `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		biz := fndbiz.NewBizUser()
		u, err := biz.Login(c, handler.DB, req.User, req.Password)
		if baseerrors.IsErrAuth(err) {
			c.Status(http.StatusUnauthorized)
			return
		} else if c.HandleError(err) {
			return
		}

		jwt, err := handler.JwtSvc.NewJWT(baseservice.JWTData{
			ID: u.ID,
		})

		if c.HandleError(err) {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"User": gin.H{
				"ID":         u.ID,
				"Name":       u.Name,
				"MiddleName": u.MiddleName,
				"LastName":   u.LastName,
				"Email":      u.Email,
				"Username":   u.Username,
			},
			"JWT": jwt,
		})
	}
}

// RequestRestoreAccount creates an account access to allow the user change the password.
func (handler httpPublic) RequestRestoreAccount() httphandler.HandlerFunc {
	type Req struct {
		Email string `binding:"required,email"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		tx := c.BeginTx(handler.DB)
		biz := fndbiz.NewBizUser()
		_, err := biz.RequestRestoreAccount(c, tx.Tx, handler.MailSvc, req.Email)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}

// RestoreAccount creates an account access to allow the user change the password.
func (handler httpPublic) RestoreAccount() httphandler.HandlerFunc {
	type Req struct {
		Key      string `binding:"required"`
		Password string `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		tx := c.BeginTx(handler.DB)
		biz := fndbiz.NewBizUser()
		_, err := biz.RestoreAccount(c, tx.Tx, handler.MailSvc, req.Key, req.Password)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}
