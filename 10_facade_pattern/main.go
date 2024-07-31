package main

import (
	"errors"
	"fmt"
)

type db struct {
	datas []string
}

func (this db) query(index int) (string,error) {
	if (index >= len(this.datas) || index <0){
		return "",errors.New("out of index")
	}
	return this.datas[index],nil
}

type gin struct {

}
func (gin) handle(database db,index netPack)(string,error) {
	return database.query(index.param)
}

type netPack struct {
	param int
	// other params,headers...
}

type net struct {

}

func (*net) send(param int) netPack {
	return netPack{param:param }
}

type api struct {
	dbInstance db
	ginInstance gin
	netInstance net
}

func createApi(data []string) api {
	return api{dbInstance:db{datas: data},ginInstance: gin{},netInstance: net{}}
}

func (this api) handle(index int) (string,error) {
	pack := this.netInstance.send(index)
	return this.ginInstance.handle(this.dbInstance,pack)
}

func main() {
	api := createApi([]string{"1","2","3"})
	fmt.Println(api.handle(2))

	_db := db{datas: []string{"1","2","3"}}
	_gin := gin{}
	_net := net{}
	pack := _net.send(2)
	fmt.Println(_gin.handle(_db,pack))
}