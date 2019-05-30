package main

import (
	"github.com/nathan-osman/go-rpigpio"
	"time"
	"fmt"
)

func main(){
	p,err := rpi.OpenPin(2,rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	//set high
	p.Write(rpi.HIGH)
	fmt.Println("start gpio 2 test")

	go func() {
		for{
			time.Sleep(time.Millisecond * 100)
			p.Write(rpi.HIGH)
			fmt.Println("on")
			time.Sleep(time.Millisecond * 100)
			p.Write(rpi.LOW)
			fmt.Println("off")
		}
	}()

	time.Sleep(time.Hour * 2)
}

