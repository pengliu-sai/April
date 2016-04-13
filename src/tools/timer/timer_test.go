package timer

import (
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	c := make(chan int)
	count := 0

	DoTimer(int64(10), func() {
		log.Println(time.Now().Unix())
		count += 1
		if count > 10 {
			c <- 999
		}
	})

	log.Println(<-c)
}
