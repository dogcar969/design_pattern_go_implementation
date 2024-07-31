package main

import (
	"fmt"
	"os"
)

var isclean = true

type closeMethod interface {
	close(file *os.File) error
}

type justClose struct {
}

func (this justClose) close(file *os.File) error {
	return file.Close()
}

type cleanAndClose struct {
	
}

func (this cleanAndClose) close(file *os.File) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}
	return file.Close()
}

func fileClose(file *os.File,method closeMethod) {
	method.close(file)
}

func main() {
	fileInfo, err := os.Stat("E:/go_design_partern/21_strategy_pattern/test.txt")
	if err != nil {
		fmt.Println(err)
		return
}
	file, err := os.OpenFile("E:/go_design_partern/21_strategy_pattern/test.txt",os.O_RDWR,0)
	if err != nil {
			fmt.Println(err)
			return
	}
	
	data := make([]byte, fileInfo.Size())
	count, err := file.Read(data)
	if err != nil {
			return
	}
	fmt.Println("字符串数据：", string(data)) // Hello world
	fmt.Println("所获取字节的长度：", count)     // 11
	if !isclean {
		fileClose(file,justClose{})
	} else {
		fileClose(file,cleanAndClose{})
	}
}