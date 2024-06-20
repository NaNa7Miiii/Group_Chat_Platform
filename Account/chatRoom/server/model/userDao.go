package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"encoding/json"
	"chatRoom/common/message"
)

// initialize an global UserDao object when the server starts
var (
	MyUserDao *UserDao
)


// define an UserDao struct to enable CRUD to User struct
type UserDao struct {
	pool *redis.Pool
}

// use a factory pattern to create an UserDao object
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao {
			pool: pool,
	}
	return
}

// return a User object and error based on userID
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { // unable to find corresponding id in users
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// unmarshal res to a User object
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 1. login - verify a user
// 2.	if the id & pwd of a user is correct, return a user object, otherwise
//    return a corresponding error message
func (this *UserDao) Login(userId int, userPassword string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	// userId is valid, verify the password
	if user.UserPwd != userPassword {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil { // user already exists
		err = ERROR_USER_EXISTS
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("Registration error=", err)
		return
	}
	return
}