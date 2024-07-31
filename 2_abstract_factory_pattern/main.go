package main

import "fmt"

type food interface {
	eat()
}

type salad struct {
}
func (* salad) eat() {
	fmt.Println("eat salad")
}

type burger struct {
}

func (*burger) eat() {
	fmt.Println("eat burger")
}

type drink interface {
	doDrink()
}

type pepsi struct {
}

func (*pepsi ) doDrink() {
	fmt.Println("drink pepsi")
}

type coca struct {
}

func (*coca ) doDrink() {
	fmt.Println("drink coca")
}

type porlar interface {
	cook(string) (food)
	makeDrink(string)(drink)
}



type kitchen struct {
}

func (* kitchen) cook(foodType string) (_food food) {

	if foodType == "burger" {
		return &burger{}
	}else if foodType == "salad" {
		return &salad{}
	}
	return nil
}

func (* kitchen) makeDrink(drinkType string) (_drink drink) {
	return nil
}

type waterBar struct {
}

func (*waterBar) makeDrink(drinkType string) (_drink drink) {
	if drinkType == "pepsi" {
		return &pepsi{}
	} else if drinkType == "coca" {
		return &coca{}
	}
	return nil
}

func (*waterBar) cook(foodType string) (food) {
	return nil
} 

type architect struct {
	
}

func (*architect) getPorlar(porlarType string) (porlar) {
	if porlarType == "kitchen" {
		fmt.Println("construct kitchen")
		return &kitchen{}
	} else if (porlarType == "waterBar") {
		fmt.Println("construct waterBar")
		return &waterBar{}
	}
	return nil
}

func main() {
	_architect := &architect{}
	_kitchen := _architect.getPorlar("kitchen")
	_waterBar := _architect.getPorlar("waterBar")
	_salad := _kitchen.cook("salad")
	if _salad != nil {_salad.eat()}
	_burger := _kitchen.cook("burger")
	if _burger != nil {_burger.eat()}
	_pepsi := _waterBar.makeDrink("pepsi")
	if _pepsi != nil {_pepsi.doDrink()}
	_coca := _waterBar.makeDrink("coca")
	if _coca != nil {_coca.doDrink()}
}