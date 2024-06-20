package message

// define an user struct
type User struct {
	UserId int `json:"userId"` // add tag for marshal and unmarshal
	UserPwd string `json:"userPWd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`
}

