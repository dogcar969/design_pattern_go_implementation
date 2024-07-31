package main

import (
	"errors"
	"fmt"
	"strconv"
)

type iterator[T any] struct {
	array []T
	index int
}

func (iter *iterator[T]) hasNext() bool {
	if iter.index < len(iter.array) {
		return true
	} else {
		return false
	}
}

func (iter *iterator[T]) next() (int,T,error) {
	if iter.hasNext() {
		iter.index++
		return iter.index-1,iter.array[iter.index-1],nil
	}
	var zero T
	return -1,zero,errors.New("获取next错误")
}

func (iter *iterator[T]) forEach(Func func(int,T)any)error{
	for iter.hasNext() {
		index,item,err := iter.next()
		if err != nil {
			return err
		}
		Func(index,item)
	}
	iter.index = 0
	return nil
} 
func (iter *iterator[T]) _map1(Func func(int,T)any) (*iterator[any],error) {
	// 不限制输出的新数组的类型
	var newIter iterator[any]
	for iter.hasNext() {
		index,item,err := iter.next()
		if err!= nil {
			var zero iterator[any]
			return &zero,err
		}
		newIter.array = append(newIter.array, Func(index,item))
	}
	iter.index = 0
	return &newIter,nil
}

func (iter *iterator[T]) _map2(Func func(int,T)T) (*iterator[T],error) {
	// 限制新数组的类型与旧数组一样
	var newIter iterator[T]
	for iter.hasNext() {
		index,item,err := iter.next()
		if err!= nil {
			var zero iterator[T]
			return &zero,err
		}
		newIter.array = append(newIter.array, Func(index,item))
	}
	iter.index = 0
	return &newIter,nil
}


func arrayTimes2(index int, item int) int {
	return item * 2
}

func intToStr(index int,item int) any {
	return strconv.Itoa(item)
}

func arrayPrint[T any](index int,item T) any{
	fmt.Println("第" , index , "个元素是", item)
	return nil
}

func main() {
	var iter1 = iterator[int]{array: []int{1,2,3,4,5}}
	var iter2,err = iter1._map2(arrayTimes2)
	iter2.forEach(arrayPrint)
	if err != nil {
		fmt.Println(err)
	}
	iter3,err := iter2._map1(intToStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(iter3.array...)
	iter3.forEach(arrayPrint)
	
	
}