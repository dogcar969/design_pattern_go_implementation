package instance

import "fmt"

type singleton struct {
}

func (*singleton) ShowCount() {
	fmt.Println(Count)
}

func newSingleton() *singleton {
	Count += 1
	return &singleton{}
}

var Count = 0

var instance = newSingleton()

func GetInstance() *singleton { return instance }
