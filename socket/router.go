// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package socket

import (
	"github.com/astaxie/beego"
	"github.com/googollee/go-socket.io"
	"net/http"
	"time"
	"unsafe"
	"wing/logger"
	"wing/utils"
)

const (
	serverPingInterval = 30 * time.Second
	serverPingTimeout  = 60 * time.Second
	maxConnectCount    = 200000
)

var clientsPool *sockeClientsPool  // socket clients pool
var socketCtrller SocketController // socket controller

// sockeClientPool client connections pool
type sockeClientsPool struct {
	req2uid map[uintptr]string         // request pointer to user uuid
	uid2req map[string]uintptr         // user uuid to request pointer
	clients map[string]socketio.Socket // user uuid to socketio.Socket
}

// SocketController socket controller interface
type SocketController interface {
	OnGenEventsMap() map[string]interface{}
	OnAuthentication(req *http.Request) bool
	OnConnected(so socketio.Socket)
	OnDisconnected(so socketio.Socket)
	OnNotify(so socketio.Socket) error
	OnError(so socketio.Socket, err error)
}

// Init init socket.io handler
func Init(ctrller SocketController) {
	if ctrller == nil {
		logger.E("Invalid socket controller!")
		panic("Invalid socket controller!")
	}

	clientsPool = &sockeClientsPool{}
	clientsPool.req2uid = make(map[uintptr]string)
	clientsPool.uid2req = make(map[string]uintptr)
	clientsPool.clients = make(map[string]socketio.Socket)
	socketCtrller = ctrller

	beego.Handler("/socket.io/", socketHandler())
	logger.I("Inited app socket.io handler")
}

// Notify request notify target client
func Notify(uuid string) error {
	logger.I("Request notify client:", uuid)
	if clientsPool == nil || socketCtrller == nil {
		logger.E("SocketController have not inited!")
		return utils.ErrUnperparedState
	}

	target := clientsPool.clients[uuid]
	return socketCtrller.OnNotify(target)
}

// socketHandler return a valid http handler for socket.io
func socketHandler() http.Handler {
	logger.I("Create socket handler...")
	server, err := socketio.NewServer(nil)
	if err != nil {
		logger.E(err.Error())
		panic(err)
	}

	// set socket.io ping interval and timeout
	logger.I("Set socket ping-pong and timeout")
	server.SetPingInterval(serverPingInterval)
	server.SetPingTimeout(serverPingTimeout)

	// set max connection count
	server.SetMaxConnection(maxConnectCount)

	// set auth middleware for socket.io connection
	server.SetAllowRequest(func(req *http.Request) error {
		// prase url params
		if err := req.ParseForm(); err != nil {
			logger.E("Parse url err:", err)
			return utils.ErrInvalidData
		}

		if !socketCtrller.OnAuthentication(req) {
			logger.E("Auth client error")
			return utils.ErrAuthDenied
		}

		uuid := req.Form.Get("uuid")
		reqptr := uintptr(unsafe.Pointer(req))
		logger.I("Bind data:{uuid:", uuid, "request:", reqptr, "}")

		clientsPool.uid2req[uuid] = reqptr
		clientsPool.req2uid[reqptr] = uuid
		client := clientsPool.clients[uuid]
		if client != nil {
			client.Disconnect()
		}
		clientsPool.clients[uuid] = nil
		return nil
	})

	// set connection event
	server.On("connection", func(so socketio.Socket) {
		socketCtrller.OnConnected(so)
	})

	// set disconnection event
	server.On("disconnection", func(so socketio.Socket) {
		socketCtrller.OnDisconnected(so)
	})

	// set error event
	server.On("error", func(so socketio.Socket, err error) {
		socketCtrller.OnError(so, err)
	})

	// set custom events
	events := socketCtrller.OnGenEventsMap()
	if events != nil {
		for evt, callback := range events {
			server.On(evt, callback)
		}
	}

	logger.I("Inited socket.io handler")
	return server
}
