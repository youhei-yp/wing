// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package main

import (
	"fmt"
	"github.com/youhei-yp/wing/secure"
	"os"
)

func main() {
	uuidInt64 := secure.GenUUID()
	fmt.Println("uuid int64:", uuidInt64)

	uuidString := secure.GenUUIDString()
	fmt.Println("uuid string:", uuidString)

	code := secure.GenCode()
	fmt.Println("code:", code)

	salt, _ := secure.GenSalt()
	fmt.Println("salt:", salt)
	os.Exit(1)
}
