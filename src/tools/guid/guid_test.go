package guid

import (
	"fmt"
	"sync"
	"testing"
)

var mx sync.Mutex
var test map[uint64]int

func TestNewID(t *testing.T) {
	guid := NewGuid()

	test = make(map[uint64]int)
	for i := 0; i < 10000; i++ {
		go func() {
			id := guid.NewID(1)
			addID(id, t)
		}()
	}
}

func addID(id uint64, t *testing.T) {
	mx.Lock()
	defer mx.Unlock()

	if _, exists := test[id]; exists {
		fmt.Println("what?", id)
		t.FailNow()
	}
	test[id] = 0
}
