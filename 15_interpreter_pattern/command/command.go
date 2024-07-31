package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Param struct {
	NickName string // --time -> -t
	Value string // 如果在设置的时候写了这个就是初始值
	HelpText string
}

type Command struct {
	Name     string
	params   map[string]Param
	HelpText string
	Func func(map[string]string)any // 会在放入参数之前将params修改为键值形式
}

var Commands map[string]Command = make(map[string]Command)

func CreateCommand(name string, _params map[string]Param,HelpText string,Func func(map[string]string)any) (Command,error) {
	for name:= range _params{
		if name == "" || name == "exit" {
			return Command{},errors.New("参数名不能为空或 exit")
		}
	}
	command := Command{Name: name,params: _params,HelpText: HelpText,Func: Func}
	return command,nil
}

func GetCommand(name string) Command  {
	for _, command := range Commands {
		if command.Name == name {
			return command
		}
	}
	return Command{Name: name} // 如果找不到就新建一个
}

func UseCommand(command Command) {
	_,ok := Commands[command.Name]
	if ok {
		fmt.Println("已经有该命令，如果要修改命令，请使用SaveCommand")
		return
	}
	Commands[command.Name] = command
}

func SaveCommand(command Command) {
	_,ok := Commands[command.Name]
	if !ok {
		fmt.Println("没有该命令，如果要加入命令，使用UseCommand")
		return
	}
	Commands[command.Name] = command
}

func DeleteCommand(name string) {
	delete(Commands,name)
}

func nickNameMatch(key string,params map[string]Param) string {
	for name,param := range params {
		if param.NickName == key {
			return name
		}
	}
	return ""
}

var Out any

func receiveCommand(text string) error {
	// 解析接收到的命令，如果成功匹配就执行
	if text == "" {
		return errors.New("无法解析空字符串")
	}
	segments := strings.Split(text, " ")
	name := segments[0]
	command,ok := Commands[name]
	if !ok {
		return errors.New("没有匹配的命令名")
	}
	keyFlag := true // 判断现在解析的是什么
	key:= ""
	for _,param := range segments[1:] {
		if keyFlag {
			keyFlag = false
			if len(param)>2 {
				if param[:2] == "--" {
					key = param[2:]
					_,ok = command.params[key]
					if !ok {
						return errors.New("参数名匹配错误："+key)
					}
				} else if param[:1] == "-"{
					key = param[1:]
					key = nickNameMatch(key,command.params)
					if key == "" {
						return errors.New("参数名匹配错误："+param[1:])
					}
				}
			} else {
				return errors.New("参数格式错误")
			}
		} else {
			keyFlag = true
			if entry,ok := command.params[key]; ok {
				entry.Value = param
				command.params[key] = entry
			}
		}
	}
	_params := map[string]string{}
	for k,v := range command.params { // 不与上面的重复
		_params[k] = v.Value
	}

	Out = command.Func(_params)
	fmt.Println(Out)
	return nil
}

func commandStringify(name string,command Command) string {
	helpText := "usage: " + name
	for paramName,param := range command.params {
		helpText +=" --" +paramName
		if param.NickName != ""{
			helpText += "(-" +param.NickName + ")"
		}
	}
	helpText += "\n" + command.HelpText + "\n"
	for paramName,param := range command.params {
		if param.HelpText != "" {
			helpText += paramName + " " + param.HelpText + "\n"
		}
	}
	helpText += "\n"
	return helpText
}

func Run(){
	go func() {
		reader := bufio.NewReader(os.Stdin)
		commandText := ""
		data,_,err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		commandText = string(data)
	for commandText != "exit" {
		if commandText == "help" {
			for name,command := range Commands {
				fmt.Println(commandStringify(name,command))
			} 
			data,_,err := reader.ReadLine()
			if err != nil {
				fmt.Println(err)
			}
			commandText = string(data)
			continue
		}
		err := receiveCommand(commandText)
		if err != nil {
			fmt.Println("command err:",err)
		}
		data,_,err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		commandText = string(data)
	}
	}()
}
