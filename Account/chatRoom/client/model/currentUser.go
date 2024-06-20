package model

import (
	"net"
	"chatRoom/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}