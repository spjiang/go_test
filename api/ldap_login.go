
package api

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap"
	"github.com/kataras/iris/context"
	"ldap/conf"
)

func ActionLdapLogin(ctx context.Context) {
	params := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := ctx.ReadJSON(&params); err != nil {
		Err(ctx, ERROR_PARAM, "Json parse error.", err)
		return
	}

	conn, err := ldap.DialTLS("tcp", conf.Conf.Ldap.Host + ":" + conf.Conf.Ldap.Port, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		Err(ctx, ERROR_LDAP, "Ldap server disconnect.", err)
		return
	}
	defer conn.Close()

	err = conn.Bind(params.Username, params.Password)
	if err != nil {
		Err(ctx, ERROR_PASSWORD, "Password error.", err)
		return
	}

	sql := ldap.NewSearchRequest(conf.Conf.Ldap.Base,
		ldap.ScopeWholeSubtree,
		ldap.DerefAlways,
		0,
		0,
		false,
		fmt.Sprintf("(sAMAccountName=%s)", params.Username),
		[]string{"sAMAccountName", "displayName", "mail", "mobile", "employeeID", "givenName"},
		nil)

	var cur *ldap.SearchResult

	if cur, err = conn.Search(sql); err != nil {
		Err(ctx, ERROR_LDAP, "Ldap server search failed.", err)
		return
	}

	if len(cur.Entries) == 0 {
		Err(ctx, ERROR_NOUSER, "Not found user.", nil)
		return
	}

	var result = struct {
		Name               string `json:"name"`
		Account            string `json:"account"`
		Email              string `json:"email"`
		Phone              string `json:"phone"`
		EmployeeId         string `json:"employeeId"`
	}{
		Name:               cur.Entries[0].GetAttributeValue("givenName"),
		Account:            cur.Entries[0].GetAttributeValue("sAMAccountName"),
		Email:              cur.Entries[0].GetAttributeValue("mail"),
		Phone:              cur.Entries[0].GetAttributeValue("mobile"),
		EmployeeId:         cur.Entries[0].GetAttributeValue("employeeID"),
	}

	Suc(ctx, &result)
}
