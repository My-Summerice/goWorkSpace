package main

import ( 
 	"fmt" 
 	"time"
)

func consumer(cname string, data chan int) { 
 	for i := range data {
 		fmt.Println("consumer-----------", cname, ":", i) 
 	} 
 }
 
func producer(pname string, data chan int) { 
 	for i := 0; i < 2; i++ { 
 		fmt.Println("producer--", pname, ":", i) 
 		data <- i 
 	}
}
 
 
 func main() { 
 
 	data := make(chan int) 

 	go producer("生产者1", data) 
 	go producer("生产者2", data) 
 	go consumer("消费者1", data) 
 	go consumer("消费者2", data) 

 	time.Sleep(2 * time.Second) 
 	close(data) 
 	time.Sleep(2 * time.Second)
 }