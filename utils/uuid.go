// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package utils

import (
	"github.com/bwmarrin/snowflake"
	"strconv"
	"wing/logger"
)

// uuidNode : generate uuid string
var uuidNode *snowflake.Node

// init uuid generater
func init() {
	if uuidNode == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			logger.E("Create uuid generater, err:", err.Error())
			panic(err.Error())
		}
		logger.I("Inited uuid generater:", node)
		uuidNode = node
	}
}

// ObatinUUID generate a new uuid in int64
func ObatinUUID() int64 {
	return uuidNode.Generate().Int64()
}

// ObatinUUIDString generate a new uuid in string
func ObatinUUIDString() string {
	return strconv.FormatInt(ObatinUUID(), 10)
}
