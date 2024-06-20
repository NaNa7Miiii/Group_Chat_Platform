package process

import (
	"fmt"
	"net"
	"chatRoom/common/message"
	"encoding/json"
	"chatRoom/server/utils"
	"chatRoom/server/model"
)

type UserProcessor struct {
	Conn net.Conn
	UserId int // shows who the conn belongs to
}

// notify all the other online users
func (this *UserProcessor) NotifyOtherOnlineUsers(userId int) {
	// loop through onlineUsers, individually send 'NotifyUserStatusMes' message
	for id, userProcessor := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		userProcessor.NotifyMyOnlineStatus(userId)
	}
}

func (this *UserProcessor) NotifyMyOnlineStatus(userId int) {
	// assemble NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.UserOnline

	// marshal notifyUserStatusMes
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMyOnlineStatus err=", err)
		return
	}
}



func (this *UserProcessor) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	// register through redis DB
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "Unknown Registration error"
		}
	} else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// convert data to resMes
	resMes.Data = string(data)

	// marshal resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// send data via writePkg func
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return err
}

func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
	// take mes.Data from mes and unmarshal it to LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	// 1.declare a resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2.declare a LoginResMes
	var loginResMes message.LoginResMes

	// verify the user through the redis DB
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPassword)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "Server Error"
		}
	} else {
		loginResMes.Code = 200
		// add userId of the user successfully logged in to 'this'
		this.UserId = loginMes.UserId
		// put the user logged in to userMgr
		userMgr.AddOnlineUsers(this)
		this.NotifyOtherOnlineUsers(loginMes.UserId) //?
		// put the user id to loginResMes.UserIds
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "Successful Login")
	}

	// 3.Marshal loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// 4.convert data to resMes
	resMes.Data = string(data)

	// 5.marshal resMes
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	// 6.send data via writePkg func
	tf := &utils.Transfer {
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	return err
}

