package concurrency

import (
	"time"
)

/**
* Produces numbers 1..n, for each consumer. each consumer can pause the production and ask to resume again
 */
const (
	START = iota
	PAUSE
	EXIT
)

type Producer struct {
	value int
	cmd   chan int
	done  chan struct{}
	data  chan int
}

func (p *Producer) subscribe() {
	go func() {
		for {
			select {
			case c := <-p.cmd:
				if c == START {
					go func() {
						for {
							select {
							case <-p.done:
								return
							default:
								p.data <- p.value
								p.value++
							}
						}
					}()
				}
				if c == PAUSE {
					p.done <- struct{}{}
				}
			default:
				<-time.After(100 * time.Millisecond)
			}
		}
	}()
}

type Consumer struct {
	producer *Producer
}

func (c *Consumer) stop() {
	c.producer.cmd <- PAUSE
}
func (c *Consumer) start() {
	c.producer.cmd <- START
}

func createProducer() *Producer {
	p := &Producer{
		value: 0,
		cmd:   make(chan int),
		data:  make(chan int),
		done:  make(chan struct{}),
	}
	p.subscribe()
	return p
}
