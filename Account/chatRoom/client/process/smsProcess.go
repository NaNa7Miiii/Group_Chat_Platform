package process

import (
	"fmt"
	"chatRoom/common/message"
	"encoding/json"
	"chatRoom/client/utils"
)

type SmsProcess struct {
}

// send group messages
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 1. create a message
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2. create a SmsMes object
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	// 3. marshal smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail = ", err.Error())
		return
	}
	mes.Data = string(data)

	// 4. marshal mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail = ", err.Error())
		return
	}

	// 5. send mes to server
	tf := &utils.Transfer {
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}