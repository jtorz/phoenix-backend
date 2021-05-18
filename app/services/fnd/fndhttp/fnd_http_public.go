package fndhttp

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/authorization"
)

// HttpPublic http handler component.
type HttpPublic struct {
	DB *sql.DB
}

// Login http handler.
func (handler HttpPublic) Login(jwtSvc authorization.JWTService) httphandler.HandlerFunc {
	type request struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := request{}
		if c.BindJSON(&req) {
			return
		}

		/*ctx := handler.NewTransversalCxtAnonym(handler.DB)
		biz := fndbusiness.NewBizUser(ctx)
		u, err := biz.Login(req.User, req.Password)
		if errors.Is(err, fnderrors.ErrAuth) {
			c.Status(http.StatusForbidden)
			return
		} else if c.UnexpectedError(err) {
			return
		}

		auth := authorization.NewAuthService(c, ctx.Exe)
		jwt, err := auth.GetJWT(*u)
		if c.UnexpectedError(err) {
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
		})*/
		jwt, err := jwtSvc.NewJWT(authorization.AuthUser{
			ID: "TODO:",
		})

		if c.HandleError(err) {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"id":         "u.ID",
				"name":       "u.Name",
				"middleName": "u.MiddleName",
				"lastName":   "u.LastName",
				"email":      "u.Email",
				"username":   "u.Username",
			},
			"jwt": jwt,
		})
	}
}
