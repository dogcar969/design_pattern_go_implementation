package main

import "fmt"

type Food interface {
	freebie() Freebie
	Price() float32
}

type Fries struct {
}

func (*Fries) freebie() Freebie {
	return &Ketchup{}
}

func (*Fries) Price() float32 {
	return 4.5
}



type chickenNuggets struct {
}

func (*chickenNuggets) freebie() Freebie{
  return &SweetSourSauce{}
}

func (*chickenNuggets) Price() float32{
  return 8.2
}



type Freebie interface {
	name() string
}

type Ketchup struct {
}

func (*Ketchup) name() string {
	return "Ketchup"
}

type SweetSourSauce struct {

}

func (*SweetSourSauce) name() string {
	return "SweetSourSauce"
}

type Cook struct {
}

type meal struct {
	Foods []Food
}

func (this *meal) totalPrice() float32 {
	var total float32 = 0.0
	for _,food := range this.Foods{
		total += food.Price()
	}
	return total
}

func (this *meal) addFood(food Food) {
	this.Foods = append(this.Foods,food )
}

func (*Cook) cook(orders []string) meal {
	foods := meal{}
	if orders == nil {
		return foods
	}
	for _,order := range orders {
		if (order == "Fries") {
			foods.addFood(&Fries{})
		}else if(order == "chickenNuggets") {
			foods.addFood(&chickenNuggets{})
		}else {
			continue
		}
	}
	return foods
}

func main() {
	cook := &Cook{}
	_meal := cook.cook([]string{"Fries","chickenNuggets","Fries"})
	for _,food := range _meal.Foods{
		fmt.Println(food.freebie().name())
	}
	fmt.Println(_meal.totalPrice())

}