package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Lottery struct {
	users []*User
	time uint
	isRun bool
	Period uint // should be bigger than zero
}

func (lottery *Lottery) draw() {
	winnerIndex := rand.Int() % len(lottery.users)
	winnerName := lottery.users[winnerIndex].Name
	for index,user := range lottery.users {
		if index == winnerIndex {
			user.receive()
		} else {
			user.getInfo(winnerName)
		}
	}
	lottery.time = lottery.Period
}

func (lottery *Lottery) run() {
	if lottery.isRun {
		return
	}
	lottery.isRun = true
	go func() {
		for {
			if lottery.time == 0 {
				lottery.draw()
			} else {
				// 模拟状态变更
				lottery.time -= 1
				time.Sleep(time.Second)
			}
		}
	}()
}

func (lottery *Lottery) showRecord() {
	for _,user := range lottery.users {
		fmt.Println(user.Name,":",user.winTime)
	}
}

type User struct {
	Name string
	winTime int
}

func (user *User) receive() {
	user.winTime +=1
}

func (user *User) getInfo(winner string) {
	fmt.Println("to dear player",user.Name,",",winner,"won.","Hope you are the luck one next time.")
}



func main() {
	lottery := Lottery{users: []*User{{Name: "Alan"},{Name:"Bob"},{Name: "Cindy"}},Period: 5}
	lottery.run()
	time.Sleep(time.Second * 60)
	lottery.showRecord()
}