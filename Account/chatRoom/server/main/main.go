package main

import (
	"fmt"
	"net"
	"time"
	"chatRoom/server/model"
)

// handle communications with client
func processes(conn net.Conn) {
	defer conn.Close()
	processor := &Processor {
		Conn: conn,
	}
	err := processor.processMain()
	if err != nil {
		fmt.Println("communication goroutine err=", err)
		return
	}

}

// initialize an UserDao object using the pool defined in redis
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// initialize a redis pool
	initPool("localhost: 6379", 16, 0, 300 * time.Second)
	initUserDao()

	fmt.Println("Server is listening on port 8889...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	for {
		fmt.Println("Waiting for client to connect...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//once the connection succeed, initiate a goroutine to communicate with client
		go processes(conn)
	}
}