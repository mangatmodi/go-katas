package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	p := createProducer()
	c := &Consumer{p}
	c.start()
	pauseTick := time.After(5 * time.Second)
	startTick := time.After(15 * time.Second)
	go func() {
		for {
			select {
			case data := <-c.producer.data:
				fmt.Printf("Consumer: %d\n", data)
			case <-pauseTick:
				c.stop()
				fmt.Println("Pausing")
			case <-startTick:
				c.start()
				fmt.Println("Starting again")
			default:
				<-time.After(100 * time.Millisecond)
			}
		}
	}()

	select {}
}
