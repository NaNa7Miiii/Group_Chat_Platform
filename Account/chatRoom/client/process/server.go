package process

import (
	"fmt"
	"os"
	"net"
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
)

// show menu after successful login
func ShowMenu() {
	fmt.Println("-------------Login: XXX-------------")
	fmt.Println("----------1. Show Online Users----------")
	fmt.Println("----------2. Send Message----------")
	fmt.Println("----------3. Check Information----------")
	fmt.Println("----------4. Exit The System----------")
	fmt.Println("Please Choose from 1 to 4:")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Println("Say something to others:")
			fmt.Scanf("%s\n", &content)
			smsProcess.SendGroupMes(content)
		case 3:
		case 4:
			fmt.Println("User has logged out the system.")
			os.Exit(0)
		default:
			fmt.Println("Wrong input, please re-enter!")
	}
}

// keep connection with the client
func processServerMes(conn net.Conn) {
	// create a Transfer struct to keep reading info from client
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		switch mes.Type {
			case message.NotifyUserStatusMesType: // someone is online
				// 1.obtain NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				// 2.save the user's info and status to the client's map[int]User
				updateUserStatus(&notifyUserStatusMes)
			case message.SmsMesType: // someone sent a group message
				outputGroupMes(&mes)
			default:
				fmt.Println("Server returned an unknown message type")
		}
	}
}