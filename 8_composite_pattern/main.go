package main

import "fmt"

type army struct {
	name string
	population uint
	subUnits []army
}

func (this army) total() uint {
	if (len(this.subUnits) == 0) {
		return this.population
	}
	
	var totalNum uint
	for _,subUnit := range this.subUnits{
		subUnit.population = subUnit.total()
		totalNum += subUnit.population
	}
	fmt.Println("population",totalNum)
	return totalNum
}

func main() {
	// 一级
	var marine army
	// 二级
	for range []int{0,1,2,3,4} {
		subUnit := army{}
		for i := range []int{0,1,2,3,4} {
			subUnit.subUnits = append(subUnit.subUnits, army{population: uint(100 * i)})
		}
		marine.subUnits = append(marine.subUnits, subUnit)
	}
	marine.population = marine.total()
	fmt.Println(marine.population)
	
}