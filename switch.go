package main

import "fmt"
import "time"

func main(){

	t:= time.Now()
	switch {
		case t.Hour() < 12 : 
			fmt.Println(t.Hour()," time current is now")
		default : 
			fmt.Println("not now")
	}
}