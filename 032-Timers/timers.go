package main

import ("time";"fmt")

func main() {
	
	timer1 := time.NewTimer(time.Second)

	<-timer1.C
	fmt.Println("timer1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 expired")
	}()

	stop2 := timer2.Stop()
	if (stop2) {
		fmt.Println("timer2 stopped")
	}


	<-time.NewTimer(time.Second*2).C
}