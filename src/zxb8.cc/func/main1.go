package main

import "fmt"

func testSlice0(){
	var a[]int
	if a == nil{
		fmt.Printf("a is nil\n")
	}else{
		fmt.Printf("a=%v\n",a)
	}
}

func testSlice2(){
	c := []int{6,7,8}
	fmt.Println(c)
}

func testSlice3(){
	a :=[...]string{"a","b","c"}
	fmt.Print(a)
}

func testSlice4(){
	var a[]int
	fmt.Printf("%p,len:%d,cap:%d",a,len(a),cap(a))
	if a == nil{
		fmt.Printf("a is nil\n")
	}
	a = append(a,100)
	a = append(a,200)
	fmt.Printf("%p,len:%d,cap:%d\n",a,len(a),cap(a))
}


func testSlice5(){
	var a []int = []int{1,3,4}
	var b []int = []int{5,6,7}
	a = append(a,23,4,55)
	fmt.Println(a)
	a =append(a,b...)
	fmt.Printf("a=%v\n",a)
}

func sumArray11(a []int) int{
	var  sum int  = 0
	for _,v := range a{
		sum = sum + v
	}
	return sum
}

func testSumArray(){
	var a[10]int = [10]int{1,3,3,4,5,5,8}
	sum := sumArray11(a[:])
	fmt.Println("sum:",sum)

	var b [3]int = [3]int{10,20,30}
	sum = sumArray11(b[:])
	fmt.Println("sum:",sum)
}

func modifySlice(a []int){
	a[0] = 1000
}

func testModifySlice(){
	var a[3]int = [3]int{1,2,3}
	modifySlice(a[:])
	fmt.Println(a)
}

func testLearn(){
	var sa = make([]string,5,10)
	for i:=0;i<10;i++{
		sa = append(sa,fmt.Sprintf("%v",i))
	}
	fmt.Println(sa)
	fmt.Println(len(sa),cap(sa))
}

func main(){
	testSlice0()
	testSlice2()
	testSlice3()
	testSlice4()
	testSlice5()
	testSumArray()
	testModifySlice()
	testLearn()
}