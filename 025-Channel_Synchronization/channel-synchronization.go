package main

import ( "fmt" ;"time")

func worker(done chan bool) {

	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true

}

func main() {
	done := make(chan bool,1)
	go worker(done)
	time.Sleep(time.Second/2)
	go worker(done)
	<-done

	fmt.Println("Bottom of main")
}