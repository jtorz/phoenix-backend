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
		Password   string `binding:"required"`
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
		biz := fndbiz.NewBizUser(tx.Tx)
		err := biz.New(c, handler.MailSvc, &u)
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

		biz := fndbiz.NewBizUser(handler.DB)
		u, err := biz.Login(c, req.User, req.Password)
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
			"user": gin.H{
				"id":         u.ID,
				"name":       u.Name,
				"middleName": u.MiddleName,
				"lastName":   u.LastName,
				"email":      u.Email,
				"username":   u.Username,
			},
			"jwt": jwt,
		})
	}
}

// RequestRestore creates an account access to allow the user change the password.
func (handler httpPublic) RequestRestore() httphandler.HandlerFunc {
	return func(c *httphandler.Context) {
		req := struct {
			Email string `binding:"required,email"`
		}{}
		if c.BindJSON(&req) {
			return
		}
		tx := c.BeginTx(handler.DB)
		biz := fndbiz.NewBizUser(tx.Tx)
		_, err := biz.RequestRestore(c, handler.MailSvc, req.Email)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}
