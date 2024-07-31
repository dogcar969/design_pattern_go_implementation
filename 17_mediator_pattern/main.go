package main

import (
	"fmt"
	"mediator/dock"
)

const DIRECT = false

func main() {
	navigator := dock.GetNavigator()

	A := dock.NewDock("A",20)
	navigator.Add(A)
	B := dock.NewDock("B",30)
	navigator.Add(B)
	C := dock.NewDock("C",25)
	navigator.Add(C)
	D := dock.NewDock("D",40)
	navigator.Add(D)
	E := dock.NewDock("E",10)
	navigator.Add(E)

	// 直接传
	if DIRECT {
		A.Transmit(20,E)
		B.Transmit(10,C)
		C.Transmit(15,E)
		D.Transmit(10,C)
		navigator.ShowDock()
		fmt.Println(dock.Time)
	} else {
		// 用导航
		A.TransmitUsingNavigator(20,"E")
		B.TransmitUsingNavigator(10,"c")
		C.TransmitUsingNavigator(15,"E")
		D.TransmitUsingNavigator(10,"C")
		navigator.ShowDock()
		fmt.Println(dock.Time)
	}
}