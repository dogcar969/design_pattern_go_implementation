package main

import "fmt"


type sqlite struct {

}

func (sqlite) query() {
	fmt.Println("sqlite query")
}

type mysql struct {

}

func (mysql)query() {
	fmt.Println("mysql query")
}

type sqliteConnector interface {
	newAndConnect(sqlite) // 新建且连接
}

type dbConenctor interface {
	Connect(mysql)
}

type dbAdapter struct {

}

func (*dbAdapter) newAndConnect(db sqlite) {
	fmt.Println("new and connect")
	db.query()
}

func (*dbAdapter) Connect(db mysql) {
	fmt.Println("connect")
	db.query()
}

func (adapter *dbAdapter) AnyDBConnect(db any) {
	res1,isSqlite := db.(sqlite)
	if isSqlite {
		adapter.newAndConnect(res1)
		return
	} 
	res2,isMysql := db.(mysql)
	if isMysql {
		adapter.Connect(res2)
		return
	} 
		fmt.Println("db not support")
}

func main() {
	var adapter dbAdapter
	adapter.AnyDBConnect(mysql{})
	adapter.AnyDBConnect(sqlite{})
	adapter.AnyDBConnect(1)
}