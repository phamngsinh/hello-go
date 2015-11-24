package main

import "fmt"

func main() {

	fmt.Println("simple loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i=i+1
	}
	fmt.Println("loop like c++")
	for j:=0;j<=10;j++{
		fmt.Println(j)

	}
	fmt.Println("shor loop")
	t:=0
	for {
		fmt.Println(t)
		t=t+2
		if(t==10){
			break
		}
	}
	fmt.Println("array loop")
	arr:=[]int{1,2,3,4,556,6,7,78,8,54,67,8}
	for key,value :=range arr{
		fmt.Println("key:",key,";value:",value)
	}
	fmt.Println("slice loop")
	slice := map[string]int{
		"ti": 90000,
		"teo": 20000,
	}
	for key,value := range slice{
		fmt.Println("key:",key,"value:",value)
	}


}