package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap"
)

type LDAPConfig struct {
	Addr         string
	BindUserName string
	BindPassword string
	SearchDN     string
}

type LDAPService struct {
	Conn     *ldap.Conn
	Config   LDAPConfig
}


func NewLDAPService(config LDAPConfig) (*LDAPService, error) {
	conn, err := ldap.Dial("tcp", config.Addr)
	if err != nil {
		return nil, err
	}

	// NOTE(chenjun): 暂时先不skip verify
	// err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
	// if err != nil {
	//  return nil, err
	// }

	err = conn.Bind(config.BindUserName, config.BindPassword)
	if err != nil {
		return nil, err
	}
	return &LDAPService{Conn: conn, Config: config}, nil
}

// Login 登录
func (l *LDAPService) Login(userName, password string) (bool, error) {
	searchRequest := ldap.NewSearchRequest(
		l.Config.SearchDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=inetOrgPerson)(mail=%s))", userName),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Conn.Search(searchRequest)
	if err != nil {
		return false, err
	}

	if len(sr.Entries) != 1 {
		return false, fmt.Errorf("User does not exist or too many entries returned")
	}

	userDN := sr.Entries[0].DN
	err = l.Conn.Bind(userDN, password)
	if err != nil {
		return false, err
	}

	err = l.Conn.Bind(l.Config.BindUserName, l.Config.BindPassword)
	if err != nil {
		return false, nil
	}

	return true, nil
}
