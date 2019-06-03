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
	secretkey := secure.GenAESKey()
	fmt.Println(secretkey)
	os.Exit(1)
}
