package main

import ( 
	"fmt"
	"time"
)

func webServerWorker(workerId int, msg chan int )  {
	for res := range msg {
		fmt.Println("Workerid:",workerId, " Mensagem processada: ", res)
		time.Sleep(time.Second)
	}
}

func execThread()  {
	msg := make(chan int)

	go webServerWorker(1,msg)
	go webServerWorker(2,msg)
	go webServerWorker(3,msg)
	go webServerWorker(4,msg)
	go webServerWorker(5,msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}
}