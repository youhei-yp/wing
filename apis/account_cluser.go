// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package apis

import (
	"fmt"
	"github.com/astaxie/beego"
)

// getCluserHost prase cluser api url by name, router and version,
// the cluser name must config in ./conf/app.conf file
func getCluserHost(cluser, router string, ver ...string) string {
	host, vx := beego.AppConfig.String(cluser), "v1"
	if len(ver) > 0 && ver[0] != "" {
		vx = ver[0]
	}
	return fmt.Sprintf("%s/%s/%s", host, vx, router)
}

// ApiAccDelete return the API url for delete account
func ApiAccDelete(ver ...string) string {
	return getCluserHost("wgdasurl", "delete", ver...)
}

// ApiAccLogin return the API url for account login
func ApiAccLogin(ver ...string) string {
	return getCluserHost("wgdasurl", "login", ver...)
}

// ApiAccOverPwd return the API url for override account password
func ApiAccOverPwd(ver ...string) string {
	return getCluserHost("wgdasurl", "over/pwd", ver...)
}

// ApiAccProfile return the API url for get account profile
func ApiAccProfile(ver ...string) string {
	return getCluserHost("wgdasurl", "profile", ver...)
}

// ApiAccPubKey return the API url for get account RSA public key
func ApiAccPubKey(ver ...string) string {
	return getCluserHost("wgdasurl", "pubkey", ver...)
}

// ApiAccRegister return the API url for register account
func ApiAccRegister(ver ...string) string {
	return getCluserHost("wgdasurl", "register", ver...)
}

// ApiAccToken return the API url for get account token
func ApiAccToken(ver ...string) string {
	return getCluserHost("wgdasurl", "token", ver...)
}

// ApiAccUpdatePwd return the API url for update account password
func ApiAccUpdatePwd(ver ...string) string {
	return getCluserHost("wgdasurl", "update/pwd", ver...)
}

// ApiAccUpdateRole return the API url for update account role
func ApiAccUpdateRole(ver ...string) string {
	return getCluserHost("wgdasurl", "update/role", ver...)
}

// ApiAccViaToken return the API url for verify account request token
func ApiAccViaToken(ver ...string) string {
	return getCluserHost("wgdasurl", "via", ver...)
}
