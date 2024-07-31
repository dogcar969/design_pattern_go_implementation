package main

import (
	"fmt"
	"interpreter/command"
	"strconv"
)

// 做一个简单的命令行解析器



func Add(params map[string]string) any {
	// --left -l 左值
	// --right -r 右值
	l,err := strconv.Atoi(params["left"])
	if err != nil {
		fmt.Println("string parse err:",err)
	}
	r,err := strconv.Atoi(params["right"])
	if err != nil {
		fmt.Println("string parse err:",err)
	}
	return l+r
}

func main() {
	var AddParams = map[string]command.Param{}
	AddParams["left"] = command.Param{NickName:"l",HelpText: "加法的左值"}
	AddParams["right"] = command.Param{NickName:"r", HelpText: "加法的右值"}
	AddCommand,err := command.CreateCommand("Add",AddParams,"Add two number",Add)
	if err != nil {
		fmt.Println("command create err:",err)
	}
	command.UseCommand(AddCommand)
	command.Run()
	for {
		
	}
}