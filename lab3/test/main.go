package main
import ("fmt"; "time")
 
var ch1 chan int = make(chan int);
 
func foo(ch chan int) {
	for i := 0; i<100; i++ { 
        	fmt.Println(1)
    	}
}
 
func bar(ch chan int) {
	for i := 0; i<100; i++ { 
        	fmt.Println(2)
    	}
}
 
func main(){
    go bar(ch1)
    go foo(ch1)
    time.Sleep(1000)
}
