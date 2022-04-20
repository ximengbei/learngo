package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播的channel
	Message chan string
}

//创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

//监听Message广播消息channel的groutine，一旦有消息就发送给全部在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

//广播消息方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

//处理当前链接的业务
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")

	user := NewUser(conn, this)
	user.Online()

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err : ", err)
				return
			}
			msg := string(buf[:n-1])
			//处理针对用户的消息
			user.DoMessage(msg)
		}
	}()

	//当前handler阻塞
	select {}
}

//启动服务器接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err : ", err)
		return
	}
	//close listen socket,一旦结束就关闭
	defer listener.Close()

	//启动监听message的groutine
	go this.ListenMessager()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err : ", err)
			continue
		}
		//do handler
		go this.Handler(conn)
	}
}
