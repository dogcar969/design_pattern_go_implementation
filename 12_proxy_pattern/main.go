package main

import (
	"fmt"
	"proxy/context"
	"reflect"
)

// 引用

var time = 0

type RealImage struct {
	fileName string
}

func (image *RealImage) display() {
	fmt.Println(image.fileName)
}

func (image *RealImage) loadFromDisk(name string) {
	image.fileName = name
	time += 1 
	fmt.Println("load ",name," from disk")
}

type ImageProxy struct {
	realImage RealImage
	fileName string
}

func (image *ImageProxy) display() {
	if reflect.DeepEqual(image.realImage,RealImage{}) {
		image.realImage = RealImage{}
		image.realImage.loadFromDisk(image.fileName)
		image.realImage.display()
	}else {
		image.realImage.display()
	}
}

// 中间件
func procedure1(task *context.Task) {
	fmt.Println("Procedure1 begin.")
	fmt.Println("Procedure1 end.")
}

func procedure2(task *context.Task) {
	fmt.Println("Procedure2 begin.")
	task.Next()
	fmt.Println("Procedure2 end.")
}

func procedure3(task *context.Task) {
	fmt.Println("Procedure3 begin.")
	fmt.Println("Procedure3 end.")
}


func main() {
	// 引用
	var image = ImageProxy{fileName: "icon.png"}
	fmt.Println(time)
	image.display() 
	fmt.Println(time)
	image.display() 
	fmt.Println(time)

	// 中间件
	task := context.NewTask(context.H{},procedure1,procedure2,procedure3)
	task.Do()
}