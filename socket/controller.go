// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : liuchuan, yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   liuchuan       New version
// 00002       2019/03/01   yangping       Code refactoring
// -------------------------------------------------------------------

package socket

import (
	"github.com/googollee/go-socket.io"
	"net/http"
	"unsafe"
	"wing/logger"
)

// SimpleController default socket controller
type SimpleController struct {
	SocketController
}

// OnGenEventsMap genernate socket events map
func (s *SimpleController) OnGenEventsMap() map[string]interface{} {
	return nil
}

// OnAuthentication event of authentication
func (s *SimpleController) OnAuthentication(req *http.Request) bool {
	return true
}

// OnConnected event of connected
func (s *SimpleController) OnConnected(so socketio.Socket) {
	logger.I("OnConnected, Bind client socket:", so.Id())
	reqptr := uintptr(unsafe.Pointer(so.Request()))
	uuid := clientsPool.req2uid[reqptr]
	if uuid != "" {
		clientsPool.clients[uuid] = so
	}
}

// OnDisconnected event of disconnected
func (s *SimpleController) OnDisconnected(so socketio.Socket) {
	logger.I("OnDisconnected, Unbind client socket:", so.Id())
	reqptr := uintptr(unsafe.Pointer(so.Request()))
	uuid := clientsPool.req2uid[reqptr]
	if uuid != "" {
		clientsPool.clients[uuid] = nil
	}
}

// OnNotify event of notify client
func (s *SimpleController) OnNotify(so socketio.Socket) error {
	return nil
}

// OnError event of error
func (s *SimpleController) OnError(so socketio.Socket, err error) {
	logger.E("OnError, socket:", so.Id(), " error:", err)
}
