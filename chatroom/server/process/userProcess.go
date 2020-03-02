package process2

import (
	"encoding/json"
	"fmt"
	"net"
	message "test/chatroom/common"
	"test/chatroom/server/model"
	"test/chatroom/server/utils"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

func (u *UserProcess) NotifyOthersOnlineUser(userId int) {
	// 遍历onlineUsers,然后一个一个的发送，NotifyUsersStatusMes
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}

		// 开始通知，单独写一个方法
		err := up.NotifyOthersOnline(userId)
		if err != nil {
			fmt.Println("NotifyOthersOnline err=", err)
			continue
		}
	}
}

func (u *UserProcess) NotifyOthersOnline(userId int) (err error) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyOthersOnline err=", err)
		return
	}
	return
}

// 编写一个函数serverProcessLogin函数，专门处理登录请求
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码
	// 1、先从mes 中取出 mes.Data,并直接反序列化LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("message.LoginMes 序列化错误...")
		return
	}

	// 2、定义一个返回数据结构
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 在声明一个loginResMes
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 // 500状态码，表示用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		// 这里，因为用户登录成功，我们就把该登录成功的用户放入到userMgr中
		u.UserId = loginMes.UserId

		// 增加到在线用户列表中
		userMgr.AddOnlineUsers(u)
		// 通知到其它用户
		u.NotifyOthersOnlineUser(loginMes.UserId)

		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登录成功")
	}
	//3、将loginResMes序列化
	dataRes, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println(" json.Marshal(loginResMes) fail err", err)
		return
	}
	//4、组装总的数据返回结构体
	resMes.Data = string(dataRes)

	// 5、序列化总的返回结构体
	body, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail err=", err)
	}
	// 6、发送状态
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(body)
	if err != nil {
		fmt.Println("writePkg fail err=", err)
		return
	}
	return nil
}

// 用户注册
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("message.registerMes 序列化错误...")
		return
	}

	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "服务器内部错误"
		}
	} else {
		registerResMes.Code = 200
	}

	//3、将loginResMes序列化
	dataRes, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println(" json.Marshal(loginResMes) fail err", err)
		return
	}
	//4、组装总的数据返回结构体
	resMes.Data = string(dataRes)

	// 5、序列化总的返回结构体
	body, err := json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail err=", err)
	}
	// 6、发送状态
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	err = tf.WritePkg(body)
	if err != nil {
		fmt.Println("writePkg fail err=", err)
		return
	}
	return nil
}
