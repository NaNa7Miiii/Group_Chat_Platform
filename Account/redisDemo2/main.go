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
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	// 3. Read data
	// 因为返回的r是interface{}
	// 因为name对应的值是string，因此我们需要转换
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return
	}
	r2, err := redis.Int(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return
	}


	fmt.Printf("r1=%v, r2=%v", r1, r2)
}