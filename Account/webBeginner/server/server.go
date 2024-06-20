package main
import (
	"fmt"
	"net"
	"io"
)

func process(conn net.Conn) {
	defer conn.Close()
	//循环地接受客户端发送的数据
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送消息
		//2.如果客户端没有write[发送]，那么协程就阻塞在这里
		// fmt.Println("The server is waiting for client %s\n" + conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("The server exits", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("The server starts listening...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()
	fmt.Printf("listen suc=%v\n", listen)

	//循环等待客户端来连接我
	for {
		fmt.Println("Waiting for the client to connect...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v\n", conn)
		}
		//这里准备一个协程，为客户端服务
		go process(conn)
	}
}