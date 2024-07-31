package main

import (
	"errors"
	"fmt"
	"time"
)


type DB struct {
	datas []string
}

type Connection struct {
	db *DB 
	id int
	isUsed bool
}

func (this *Connection) query(index uint) (string,error) {
	if this.isUsed {
		fmt.Println("connection ",this.id,":","",errors.New("Connection conflict"))
		return "",errors.New("Connection conflict")
	} 
	this.isUsed = true
	res := this.db.datas[index]
	time.Sleep(time.Second)
	this.isUsed = false
	fmt.Println("connection ",this.id,":",res,nil)
	return res,nil
}

type Connections struct {
	Conns []*Connection
}

func NewConnections(datas []string,connNum uint) Connections {
	conns := []*Connection{}
	for i:=0;i<int(connNum);i++ {
		conns = append(conns, &Connection{db: &DB{datas: datas},id: i})
	}
	return Connections{Conns: conns}
}

func (this *Connections) query(connIndex uint,dataIndex uint)(string,error){
	if(connIndex >= uint(len(this.Conns))) {
		return "",errors.New("conn index err")
	}
	return this.Conns[connIndex].query(dataIndex)
}

func main() {
	conns := NewConnections([]string{"1","2","3"},3)
	go conns.query(0,0)
	go func () {
		time.Sleep(time.Millisecond)
		conns.query(0,0)
	}()
	go conns.query(1,2)
	time.Sleep(1* time.Second)
	conns.query(0,1)
	for ;; {

	}
}