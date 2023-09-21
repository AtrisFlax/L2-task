package main

import (
	"fmt"
	"reflect"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	jointChannel := make(chan interface{})

	var slCases []reflect.SelectCase
	for _, channel := range channels {
		slCases = append(slCases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(channel),
			//Send: interface{}
		})
	}

	go func() {
		_, _, _ = reflect.Select(slCases)
		close(jointChannel)
	}()

	return jointChannel
}

func main() {
	var sig = func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	fmt.Printf("start at  %v", start)

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
