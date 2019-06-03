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
	prikey, pubkey, _ := secure.GenRSAKeys(1024)
	fmt.Println(prikey)
	fmt.Println(pubkey)
	os.Exit(1)
}
