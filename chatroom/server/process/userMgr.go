package process2

import "fmt"

// 因为UserMgr 实例在服务器中只有一个
// 因为在很多地方都要使用，因此我们将其定义一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成一个对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers添加
func (this *UserMgr) AddOnlineUsers(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 完成对onlineUsers添加
func (this *UserMgr) DelOnlineUsers(userId int) {
	delete(this.onlineUsers, userId)
}

// 完成查询所有onlineUsers在线用户
func (this *UserMgr) GetOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUsersById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("当前用户%d,不在线", userId)
	} else {
		err = nil
	}
	return
}
