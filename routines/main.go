package main

import (
	"fmt"
	"time"
)

func main() {
	go myProcess("A")
	go myProcess("B")

	time.Sleep(3 * time.Second)

	//Canales

	myFirstChanel := make(chan string)
	go myProcessWithChannel("C", myFirstChanel)
	resul := <-myFirstChanel
	fmt.Println(resul)
	close(myFirstChanel)

	chanelA := make(chan string)
	chanelB := make(chan string)
	go myOtherProcessWithChannel("D", chanelA)
	go myOtherProcessWithChannel("E", chanelB)
	resulA := <-chanelA
	fmt.Println(resulA)

	resulB := <-chanelB
	fmt.Println(resulB)

	close(chanelA)
	close(chanelB)

}

func myProcess(p string) {
	i := 0
	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
}

func myProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok"
}

func myOtherProcessWithChannel(p string, c chan string) {
	i := 0
	for i < 15 {
		time.Sleep(150 * time.Millisecond)
		i++
		fmt.Printf("process: %s - num: %d\n", p, i)
	}
	c <- "ok 2"
}
