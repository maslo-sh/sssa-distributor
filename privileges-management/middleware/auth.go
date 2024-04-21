package middleware

import (
	auth "github.com/korylprince/go-ad-auth"
	"privileges-management/database"
)

func LDAPAuth(username, password string) bool {
	config := database.GetActiveDirectoryAuthConfig()
	status, err := auth.Authenticate(config, username, password)

	if err != nil {
		return false
	}

	if !status {
		return false
	}

	return true
}
