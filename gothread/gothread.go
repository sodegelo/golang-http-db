package gothread

import ( 
	"fmt"
	"time"
	"math"
	"math/rand"
)

func webServerWorker(workerId int, msg chan int )  {
	for res := range msg {
		res_2 := math.Sqrt(math.Sqrt(rand.Float64()))
		fmt.Println("Workerid:",workerId, " Mensagem processada: ", res," resultado:",res_2)
		//time.Sleep(time.Millisecond * 100)
	}
}

func execThread()  {
	msg := make(chan int)

	 
	for i := 0; i < 10000; i++ {
		go webServerWorker(i,msg)
	}

	for i := 0; i < 1000000; i++ {
		msg <- i
	}
	time.Sleep(time.Second * 10)
}