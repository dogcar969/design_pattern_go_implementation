package main

import "fmt"

var time =0

var db = false

type DBData struct {
	data bool
}

func newDBData() DBData {
	time += 1 // 模拟每次在数据库查找时都需要1s
	return DBData{data: db}
}

type DBDataCache struct {
	data DBData
}

func (cache *DBDataCache) init() {
	cache.data = newDBData()
}

func (cache *DBDataCache) clone(pointer *DBData) {
	*pointer = cache.data
}

func (cache *DBDataCache) save(data bool) {
	db = data
	cache.data = newDBData() // 仅在修改后重新获取
}

func main() {
	for range []int{1,2,3,4,5}{
		fmt.Println(newDBData().data) // 每次都重新查找
	}
	db = true
	for range []int{1,2,3,4,5}{
		fmt.Println(newDBData().data) // 每次都重新查找
	}
	fmt.Println(time)
	time = 0
	cache := &DBDataCache{}
	cache.init()
	datas := []DBData{}
	for range []int{1,2,3,4,5}{
		 // 每次都重新查找
		data := &DBData{}
		cache.clone(data)
		datas = append(datas, *data)
	}
	cache.save(false)
	for range []int{1,2,3,4,5}{
		// 每次都重新查找
	 data := &DBData{}
	 cache.clone(data)
	 datas = append(datas, *data)
 }
	fmt.Println(datas)
	fmt.Println(time)
}