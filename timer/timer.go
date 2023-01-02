package main

import (
	"fmt"
	"time"
)

func main() {
	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the timer before it fires. Here’s an example of that.

	timer := time.NewTimer(time.Second * 3)

	<-timer.C
	fmt.Println("timer fired")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("timer2 fired")
	}()

	stopTimer2 := timer2.Stop()
	if stopTimer2 {
		fmt.Println("timer 2 stopped")
	}
}
