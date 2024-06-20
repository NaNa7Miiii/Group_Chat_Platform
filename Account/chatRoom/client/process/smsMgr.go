package process

import (
	"fmt"
	"chatRoom/common/message"
	"encoding/json"
)

func outputGroupMes(mes *message.Message) { // mes type is SmsMesType
	// 1. Unmarshal mes data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	// show message
	info := fmt.Sprintf("User ID:\t%d sent a group message:\t%s", smsMes.UserId,
										 smsMes.Content)
	fmt.Println(info)
	fmt.Println()

}