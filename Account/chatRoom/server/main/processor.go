package main

import (
	"fmt"
	"net"
	"chatRoom/common/message"
	"chatRoom/server/process"
	"io"
	"chatRoom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			// handle login
			up := &process.UserProcessor {
				Conn: this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			// handle register
			up := &process.UserProcessor {
				Conn: this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType:
			smsProcess := &process.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default:
			fmt.Println("DNE message type, unable to process...")
	}
	return
}

func (this *Processor) processMain() (err error) {
	for {
		tf := &utils.Transfer {
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Server exits due to client logout")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}