package main

import (
	"fmt"
	"mementoPack/memento"
	"strconv"
)

func main() {
	file := memento.File{}
	// 读写 保存测试
	file.Content.SetContent("version 0")
	file.Save()
	fmt.Println(file.Content.GetContent())
	// 数组边缘测试
	file.Undo()
	fmt.Println(file.Content.GetContent())
	file.Redo()
	fmt.Println(file.Content.GetContent())
	
	for i := 0;i<10;i++ {
		file.Content.SetContent("version "+ strconv.Itoa(i+1))
		file.Save()
	}

	
	file.Undo()
	fmt.Println(file.Content.GetContent())
	file.Undo()
	fmt.Println(file.Content.GetContent())
	file.Redo()
	fmt.Println(file.Content.GetContent())
	file.Load()
	fmt.Println(file.Content.GetContent())
}