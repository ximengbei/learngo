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
	flag       int
}

//实现创建客户端函数
func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
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

//menu函数
func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scan(&flag)

	if flag > 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入合法范围内的数字！！！")
		return false
	}
}

//阻塞的run函数
func (client *Client) run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		//根据不同的模式处理不同的业务

		switch client.flag {
		case 1:
			//公聊模式
			fmt.Println("选择公聊模式...")
			break
		case 2:
			//私聊模式
			fmt.Println("选择私聊模式...")
			break
		case 3:
			//更新用户名
			fmt.Println("选择更新用户名...")
			break
		}
	}
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
	client.run()
}
