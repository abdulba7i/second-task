package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	<-or(
		createSignal(2*time.Hour),
		createSignal(5*time.Minute),
		createSignal(1*time.Second),
		createSignal(1*time.Hour),
		createSignal(1*time.Minute),
	)
	fmt.Printf("Сompleted after %v\n", time.Since(start))

	start = time.Now()
	<-or(
		createSignal(10*time.Millisecond),
		createSignal(100*time.Millisecond),
		createSignal(time.Second),
	)
	fmt.Printf("Сompleted after  %v\n", time.Since(start))

	start = time.Now()
	<-or()
	fmt.Printf("Сompleted after  %v\n", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		closedChan := make(chan interface{})
		close(closedChan)
		return closedChan
	}

	if len(channels) == 1 {
		return channels[0]
	}

	orChan := make(chan interface{})

	go func() {
		defer close(orChan)

		var remainingChannels <-chan interface{}
		if len(channels) > 2 {
			remainingChannels = or(channels[2:]...)
		}

		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-remainingChannels:
		}
	}()

	return orChan
}

func createSignal(after time.Duration) <-chan interface{} {
	signalChan := make(chan interface{})
	go func() {
		defer close(signalChan)
		time.Sleep(after)
	}()
	return signalChan
}
