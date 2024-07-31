package main

import "fmt"

type printInt interface {
	print(string)
}

type printObj struct {

}

func (printObj) print(content string) {
	fmt.Println(content)
}

type printDecor struct {
	printInstance printInt
}

func (this printDecor) print(content string) {
	this.printInstance.print(content)
}

func (this printDecor) printWithPrefix(prefix string,content string) {
	fmt.Printf(prefix)
	this.print(content)
}

func (this printDecor) printWithhooks(content string,preFun func(string), sufFun func(string)) {
	preFun(content)
	this.print(content)
	sufFun(content)
}

func main() {
	var decor printDecor = printDecor{printObj{}}
	decor.printWithPrefix("prefix ","content")
	decor.printWithhooks("content",func(content string){fmt.Println("Before ",content)},func(content string){fmt.Println("After ",content)})
}