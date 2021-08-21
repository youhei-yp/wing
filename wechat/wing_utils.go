package wechat

// SignPacket generate wechat signature string packet
func SignPacket(method, URL, timestamp, nonce, body string) string {
	packet := ""
	packet += method + "\n"
	packet += URL + "\n"
	packet += timestamp + "\n"
	packet += nonce + "\n"
	packet += body + "\n"
	return packet
}

// AuthPacket generate wechat authorization string packet
func AuthPacket(mchid, serialno, signature, timestamp, nonce string) string {
	packet := ""
	packet += "WECHATPAY2-SHA256-RSA2048 "
	packet += "mchid=\"" + mchid + "\","
	packet += "nonce_str=\"" + nonce + "\","
	packet += "timestamp=\"" + timestamp + "\","
	packet += "serial_no=\"" + serialno + "\","
	packet += "signature=\"" + signature + "\""
	return packet
}

// NotifyPacket generate wechat signature string packet
func NotifyPacket(timestamp, nonce, body string) string {
	packet := ""
	packet += timestamp + "\n"
	packet += nonce + "\n"
	packet += body + "\n"
	return packet
}
