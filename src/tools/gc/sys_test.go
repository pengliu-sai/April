package gc

import (
	"log"
	"testing"
	"time"
)

import (
	"tools/timer"
)

func TestSysGC(t *testing.T) {
	c := make(chan int)
	count := 0

	timer.DoTimer(int64(2), func() {
		log.Println(time.Now().Unix())
		onTimer()
		count += 1
		if count > 10 {
			c <- 1
		}
	})

	<-c
}
