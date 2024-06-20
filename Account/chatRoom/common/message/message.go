package message

const (
	LoginMesType         		 = "LoginMes"
	LoginResMesType      		 = "LoginResMes"
	RegisterMesType     		 = "RegisterMes"
	RegisterResMesType   		 = "RegisterResMes"
	NotifyUserStatusMesType  = "NotifyUserStatusMes"
	SmsMesType 	 						 = "SmsMes"
)

// define several user status
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)


type Message struct{
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId int `json:"userId"`
	UserPassword string `json:"userPassword"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	//status code -> 500 = user haven't signed up; 200 = login success
	Code int `json:"code"`
	UserIds []int
	Error string `json:"error"` //error message
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	//status code -> 400 = user has account; 200 = sign up success
	Code int `json:"code"`
	Error string `json:"error"` //error message
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	UserStatus int `json:"userStatus"`
}

type SmsMes struct {
	Content string `json:"content"`
	User  //anonymous struct, inherit User
}