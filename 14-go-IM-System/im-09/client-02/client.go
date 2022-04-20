package main

import (
	"flag"
	"fmt"
	"net"
)

//定义结构体
type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
}

//实现创建客户端函数
func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
	}
	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	return client
}

var serverIp string
var serverPort int

//初始化设置默认值
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器默认IP地址(127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器默认端口号(8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	//创建client
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("连接服务器失败.....")
		return
	}
	fmt.Println("连接服务器成功.....")

	//启动客户端的业务
	select {}
}
