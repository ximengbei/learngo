package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

//创建一个用户的API

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动当前监听user channel消息的goroutine
	go user.ListenMessage()
	return user
}

//用户上线
func (this *User) Online() {
	//用户上线,将用户加入到onlinemap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

//用户下线
func (this *User) Offline() {
	//用户下线,将用户从nlinemap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this, "下线")
}

//给当前对应的客户端发送消息（打印数据到终端）
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else {
		//广播消息
		this.server.BroadCast(this, msg)
	}
}

//监听当前user channel的方法，一旦有消息就发送给服务端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
