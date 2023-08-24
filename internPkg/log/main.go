package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("test")
	// log.Panic("panic test")
	// log.Fatal("my error")
	log.Print("test print")

	date := time.Now()

	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	// file, err := os.Create("myLog.log")
	if err != nil {
		log.Panic(err)
	}

	l := log.New(file, "success: ", log.LstdFlags|log.Lshortfile)
	// l1 := log.New(os.Stdout, "success: ", log.LstdFlags|log.Lshortfile)
	l.Println("test 1")
	l.Println("test 2")
	l.Println("test 3")
	l.Println("test 4")
	l.Println("test 5")
	l.SetPrefix("error: ")
	l.Println("test 6")
	l.Println("test 7")
}
