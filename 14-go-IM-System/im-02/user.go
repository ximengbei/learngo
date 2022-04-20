package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

//创建一个用户的API

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}
	//启动当前监听user channel消息的goroutine
	go user.ListenMessage()
	return user
}

//监听当前user channel的方法，一旦有消息就发送给服务端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + ",欢迎欢迎，热烈欢迎！\n"))
	}
}
