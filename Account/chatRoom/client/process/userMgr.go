package process

import (
	"fmt"
	"chatRoom/common/message"
	"chatRoom/client/model"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser

// Reveal currently online users in client
func outputOnlineUser() {
	for id, _:= range onlineUsers{
		fmt.Println("Account id:\t is online", id)
	}
}

// help handle returned NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User {
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.UserStatus
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}