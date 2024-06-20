package process
import (
	"fmt"
	"net"
	"chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"chatRoom/client/utils"
	"os"
)

type UserProcess struct {
}

// help user register
func (this *UserProcess) Register(userId int, userPassword string,
		  userName string) (err error) {
	// 1.Connect to the server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	// 2.Prepare to send data to server through conn
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.Create a LoginMes Struct
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPassword
	registerMes.User.UserName = userName

	// 4.Encoding LoginMes
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5.Fill message.Data with data
	mes.Data = string(data)

	// 6.Encoding message
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer {
		Conn : conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Register error=", err)
		return
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//unmarshal the data portion from mes to RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("Sign up Successfully, Please Login in")
		os.Exit(0)
	} else {
		fmt.Println("registerResMes.Error")
		os.Exit(0)
	}
	return
}

// help user login
func (this *UserProcess) Login(userId int, userPassword string) (err error) {
	// 1.Connect to the server
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()

	// 2.Prepare to send data to server through conn
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.Create a LoginMes Struct
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPassword = userPassword

	// 4.Encoding LoginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 5.Fill message.Data with data
	mes.Data = string(data)

	// 6.Encoding message
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	// 7. Prevent losing packages: send data's length to the server first
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// Send the length
	n, err := conn.Write(buf[:4])
	if n != 4  || err != nil {
		fmt.Println("conn.Write(buf) failed err=", err)
		return
	}

	// Send the message itself
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) failed err=", err)
		return
	}
	tf := &utils.Transfer {
		Conn : conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	//unmarshal the data portion from mes to LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// initialize current user after a successfully login
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		// shows current online user list
		fmt.Println("Current Online Users:")
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("User ID:\t", v)
			// initialize client's onlineUsers
			user := &message.User {
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		// initiate a goroutine to keep connection with server
		// it receives and reveal any data from the server
		go processServerMes(conn)

		// show secondary menu
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}