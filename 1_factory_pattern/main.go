package main

import "fmt"

// 简单工厂模式

type food interface {
	eat()
}

type salad struct {
}

func (salad) eat() {
	fmt.Println("eat salad")
}

type burger struct {
}

func (burger) eat() {
	fmt.Println("eat burger")
}

type pie struct {
}

func (pie) eat() {
	fmt.Println("eat pie")
}

type kitchen struct {

}

func (_kitchen *kitchen) cook(Type string) (_food food){
	if Type == "salad" {
		return &salad{}
	} else if Type == "burger" {
		return &burger{}
	} else if Type == "pie" {
		return &pie{}
	}
	return nil
}

// 工厂模式

type saladFactory struct {

}

func (saladFactory) cook() salad {
	return salad{}
}

type burgerFactory struct {

}

func (burgerFactory) cook() burger {
	return burger{}
}

type pieFactory struct {

}

func (pieFactory) cook() pie {
	return pie{}
}

func main() {
	// 简单工厂模式
	_kitchen := kitchen{}
	_salad := _kitchen.cook("salad")
	if _salad!=nil {_salad.eat()}
	_burger := _kitchen.cook("burger")
	if _burger != nil {_burger.eat()}
	_pie := _kitchen.cook("pie")
	if _pie != nil {_pie.eat()}
	_food := _kitchen.cook("")
	fmt.Println(_food== nil)

	// 工厂模式
	salad_factory := saladFactory{}
	salad_factory.cook().eat()
	burger_factory := burgerFactory{}
	burger_factory.cook().eat()	
	pie_factory := pieFactory{}
	pie_factory.cook().eat()
}