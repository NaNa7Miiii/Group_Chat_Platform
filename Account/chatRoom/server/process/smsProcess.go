package process

import (
	"fmt"
	"chatRoom/common/message"
	"net"
	"encoding/json"
	"chatRoom/server/utils"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	// obtain mes content
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId { // avoid sending message to oneself
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	// create an Transfer object to send data
	tf := &utils.Transfer {
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("Message transfer err=", err)
		return
	}
}