package main
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// 通过go向redis写入数据和读取数据
	// 1. connect to redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Failed to dial err=", err)
		return
	}
	defer conn.Close()
	// 2. Write data to redis
	_, err = conn.Do("Set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	// 3. Read data
	// 因为返回的r是interface{}
	// 因为name对应的值是string，因此我们需要转换
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Set err=", err)
		return
	}

	fmt.Println("Connection succeed", r)
}