package main

import (
	"errors"
	"fmt"
)


type abstructHandler struct {
	handleFunc  func(any) (any,bool)
	nextHandler *abstructHandler
}

func (handler *abstructHandler) setNextHandler(nextHandler *abstructHandler) *abstructHandler {
	handler.nextHandler = nextHandler
	return nextHandler // 为了链式设置next
}

func (handler *abstructHandler) handle(param any) {
	any,isStop := handler.handleFunc(param)
	if !isStop {
		handler.nextHandler.handle(any)
	}
}

func NewHandler(handleFunc func(any)(any,bool)) abstructHandler {
	// 为了安全只允许使用这种方法新建abstructHandler
	return abstructHandler{handleFunc: handleFunc}
}


func main() {
	requestHandler := NewHandler(func(param any) (any,bool) { fmt.Println("request",param);
	// do something 
	res:= param
return res,false})
	responseHandler := NewHandler(func(param any) (any,bool) {
		_,ok := param.(error)
		if !ok {
			fmt.Println("response:" , param)
			return nil,true
		}
		return param,false
	})
	errorHandler := NewHandler(func (param any) (any,bool)  {
		_,ok := param.(error)
		if ok {
			fmt.Println("error:" , param)
			return nil,true
		}
		return param,false
	})
	requestHandler.setNextHandler(&responseHandler).setNextHandler(&errorHandler)
	requestHandler.handle("正常完成")
	requestHandler.handle(errors.New("出错"))
}