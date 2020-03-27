package main

type Eat struct{
	type_of_food string
	quantity_of_food int
}

type X struct{
	fruits string
	eatables Eat
}

func main(){
	var y X
	y.fruits = "apple"
	y.eatables.type_of_food = "1"
	y.eatables.quantity_of_food=1
}
