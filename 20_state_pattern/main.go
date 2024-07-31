package main

import "fmt"

type State interface {
	action()
}

type Context struct {
	vibe State
}

func (cat *Context) move() {
	cat.vibe.action()
}

type Boring struct {
}

func (boring *Boring) action() {
	fmt.Println("打豆豆")
}

type Sleepy struct {

}

func (sleepy *Sleepy) action() {
	fmt.Println("睡觉")
}

type Hungry struct {

}

func (hungry *Hungry) action() {
	fmt.Println("吃饭")
}

func main() {
	var cat Context = Context{vibe: &Hungry{}}
	cat.move()
	cat.vibe = &Sleepy{}
	cat.move()
	cat.vibe = &Boring{}
	cat.move()
}