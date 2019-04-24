package main
import (
	"fmt"
	"time"
)

func func1(a int,ch chan int){
	fmt.Println("starting")
	go func_add(a,ch)
	fmt.Println("stoping")
}

func func_add(a int,ch chan int) {
	res:=0
	for i:=0; i<=a;i++{
		res+=i
	}
	fmt.Println(res)
	time.Sleep(time.Duration(3)*time.Second)
	ch<-res
}
func main(){
	ch := make(chan int)
	a:=100
	func1(a,ch)
	fmt.Println(<-ch)
}