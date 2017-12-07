package main

import "time"

func main() {
	done := time.After(500 * time.Millisecond)
	tick := time.Tick(100 * time.Millisecond)

	for {
		select {
		case <-tick:
			print("tick")
		case <-done:
			print(".done")
			return
		default:
			print(".")
			time.Sleep(10 * time.Millisecond)
		}
	}
}