package main

import (
	"instance/instance"
)

func main() {
	_instance := instance.GetInstance()
	_instance.ShowCount()
	_instance2 := instance.GetInstance()
	_instance2.ShowCount()
}