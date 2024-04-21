package database

import (
	"fmt"
	auth "github.com/korylprince/go-ad-auth"
	"gopkg.in/ldap.v3"
	"privileges-management/config"
	"strconv"
)

func GetActiveDirectoryAuthConfig() *auth.Config {
	cfg := config.GetConfig()
	return &auth.Config{
		Server:   cfg.DomainController.Host,
		Port:     cfg.DomainController.Port,
		BaseDN:   cfg.DomainController.BaseDN,
		Security: auth.SecurityStartTLS,
	}
}

func ConnectToAD() (*ldap.Conn, error) {
	conf := config.GetConfig()
	conn, err := ldap.DialURL("ldap://" + conf.DomainController.Host + ":" + strconv.Itoa(conf.DomainController.Port))
	if err != nil {
		return nil, err
	}
	//defer conn.Close()

	// Bind with admin credentials
	err = conn.Bind(conf.DomainController.AccessCredentials.Username, conf.DomainController.AccessCredentials.Password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateUser(username, password string, conn *ldap.Conn, conf config.Config) error {
	// Create user entry
	addReq := ldap.NewAddRequest(
		fmt.Sprintf("CN=%s,%s", username, conf.DomainController.BaseDN),
		[]ldap.Control{},
	)
	addReq.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "user"})
	addReq.Attribute("cn", []string{username})
	addReq.Attribute("sAMAccountName", []string{username})
	addReq.Attribute("userPassword", []string{password})
	addReq.Attribute("givenName", []string{username})
	addReq.Attribute("sn", []string{username})

	err := conn.Add(addReq)
	if err != nil {
		return err
	}

	return nil
}

func grantAccessToResource(username, resourceDN string, conn *ldap.Conn, conf config.Config) error {

	modReq := ldap.NewModifyRequest(resourceDN, []ldap.Control{})
	modReq.Add("member", []string{fmt.Sprintf("CN=%s,%s", username, conf.DomainController.BaseDN)})

	err := conn.Modify(modReq)
	if err != nil {
		return err
	}

	return nil
}
