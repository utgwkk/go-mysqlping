package main

import "time"

func Interval(d time.Duration) (<-chan struct{}, func()) {
	ch := make(chan struct{})
	go func() {
		for {
			ch <- struct{}{}
			time.Sleep(d)
		}
	}()
	return ch, func() { close(ch) }
}
