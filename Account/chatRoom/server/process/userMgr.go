package process

import (
	"fmt"
)

var (
		userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcessor
}

// initialize userMgr
func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcessor, 1024),
	}
}

// add an online user
func (this *UserMgr) AddOnlineUsers(up *UserProcessor) {
	this.onlineUsers[up.UserId] = up
}

// delete an online user
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)
}

// search online users
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcessor {
	return this.onlineUsers
}

// returns an online user based on id
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcessor, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("User%d doesn't exist", userId)
		return
	}
	return
}