package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v3"
	"net/http"
	"privileges-management/middleware"
	"privileges-management/model/dto"
)

type LoginHandler interface {
	Login(ctx *gin.Context)
}

type LoginHandlerImpl struct {
	ldapConn *ldap.Conn
}

func NewLoginHandler(ldapConn *ldap.Conn) LoginHandler {
	return &LoginHandlerImpl{ldapConn: ldapConn}
}

func (lh *LoginHandlerImpl) Login(ctx *gin.Context) {
	var req dto.AuthPayload
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	authenticated := middleware.LDAPAuth(req.Username, req.Password)

	if !authenticated {
		ctx.String(http.StatusUnauthorized, "Active Directory authentication failed")
	}
}
