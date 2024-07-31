package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var data []int = []int{1, 2, 3}

type Command struct {
	name string
	param  any
	handle func(any) any
}

type Client struct {
	commends []Command
	history  [][]int

	taskChan chan Command
}

func (client *Client) AddCommend(commend Command) {
	client.commends = append(client.commends, commend)
}

func (client *Client) PushCommend(commend Command) {
	client.taskChan <- commend
}

// 假装他在另一个文件里无法直接获得
var singleton_client = Client{}

func (client *Client) run() {
	for commend := range client.taskChan {
		// 保存未操作时的状态
		client.history = append(client.history, data)
		// 多加个时间就是日志输出
		fmt.Println(commend.name)
		// 模拟耗时任务需要排队
		time.Sleep(time.Second*5)
		// 操作
		res := commend.handle(commend.param)
		// 打印任务队列长度
		fmt.Println("还有",len(client.taskChan),"个任务在队列中")
		// 此处可以有个res通道来输出，但是没必要
		fmt.Println(res)
	}
}

func (client *Client) start() {
	client.taskChan = make(chan Command, 10)
	go client.run()
}

func newCommand(name string,handler func(any) any,param any) Command {
	return Command{handle: handler,name: name,param: param}
}

func getDataReceiver(param any) any { return data }

func cleanDataReceiver(param any) any { data = []int{}; return nil }

func revokeReceiver(param any) any {
	client, ok := param.(*Client)
	if ok {
		if len(client.history) == 0 {
			return errors.New("no history")
		}
		data = client.history[len(client.history)-2] // 返回上一个状态,由于revoke也会加一个记录，所以返回两个
		client.history = client.history[:len(client.history)-2] // 去掉上一个状态
		return nil
	}
	return errors.New("input should be *client")
}

func getClient() *Client {
	if reflect.DeepEqual(singleton_client, Client{}) {
		getData := newCommand("getData",getDataReceiver,nil)
		cleanData := newCommand("cleanData",cleanDataReceiver,nil)
		revoke := newCommand("revoke",revokeReceiver,&singleton_client)
		commends := []Command{getData, cleanData,revoke}
		singleton_client = Client{commends: commends}
	} 

	return &singleton_client
}

func menuShow() {
	fmt.Println("1.get data")
	fmt.Println("2.clean data")
	fmt.Println("3.revoke")
}

func main() {
	client := getClient()
	client.start()
	var command string

	// use client
	menuShow()
	fmt.Scanln(&command)
	for command != "exit"{
		switch command {
			case "1":
				client.PushCommend(client.commends[0])
			case "2":
				client.PushCommend(client.commends[1])
			case "3":
				client.PushCommend(client.commends[2])
		}
		menuShow()
		fmt.Scanln(&command)
	}

	
}
