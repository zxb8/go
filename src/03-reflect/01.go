package main

import (
	"fmt"
	"reflect"
)

func main(){
	var x interface{}
	x = 3.14

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("x:type=%v,value=%v\n",t,v)

	goo :=x
	fmt.Printf("goo: type=%T,value=%v\n",goo,goo)

	x = &struct{name string}{}

	t = reflect.TypeOf(x)
	v = reflect.ValueOf(x)

	fmt.Printf("x: type=%v,value=%v\n",t,v)

	hoo := x

	fmt.Printf("hoo :type=%T,value=%v\n",hoo,hoo)
}

func printStructInfo(x interface{}){
	t := reflect.TypeOf(x)
	fmt.Println("--Kind-- -%v ---\n",t.Kind())

	if t.Kind() != reflect.Struct{
		fmt.Printf("ERR:Not a struct,expected struct value of `kind` struct")
		return
	}

	n := t.NumField()

	fmt.Printf("Struct of type '%v' has %v fields.ln",t,n)

	for i:=0;i<n;i++{
		tt :=t.Field(i)
		fmt.Printf("Field %v:name:%v,type:%v\n",i,tt.Name,tt.Type)
	}
	fmt.Println("---------")
}