package dock

import (
	"errors"
	"fmt"
)

type Dock struct {
	Name string
	capacity uint // private
	SendNavigator Navigator
}

type Navigator struct {
	docks map[string]*Dock
}

var navigator Navigator
var Time int = 0

func GetNavigator() *Navigator {
	if navigator.docks == nil {
		navigator.docks = make(map[string]*Dock)
	}
	return &navigator
}

func NewDock(Name string,capacity uint) Dock {
	return Dock{Name: Name,capacity: capacity,SendNavigator: navigator}
}

func (nav *Navigator) Add(dock Dock){
	nav.docks[dock.Name] = &dock
}

func (nav *Navigator) ShowDock() {
	for name,dock := range nav.docks {
		fmt.Println("dock " + name +": ",dock.capacity)
	}
}

func (dock *Dock) Transmit(volumn uint,to Dock) error {
	Time += 1 // 先运输
	if to.capacity<volumn {
		return errors.New("目标港口无法接收")
	} 
	dock.capacity += volumn
	to.capacity -= volumn
	return nil
}

func (nav *Navigator) send(from *Dock,volumn uint,to string) error {
	entry,ok := nav.docks[to]
	if !ok {
		from.capacity -= volumn
		return errors.New("码头检索失败")
	}
	if entry.capacity<volumn {
		return errors.New("目标港口无法接收")
	} 
	Time += 1
	from.capacity += volumn
	entry.capacity -= volumn
	nav.docks[to] = entry
	return nil
}

func (dock *Dock) TransmitUsingNavigator(volumn uint,to string) error {
	err := dock.SendNavigator.send(dock,volumn,to)
	return err
}