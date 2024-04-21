package handlers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v3"
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

}
