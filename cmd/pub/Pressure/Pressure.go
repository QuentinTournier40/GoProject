package main

import (
	"goproject/cmd/pub/Subscribers"
)

func main() {
	Subscribers.RunSubscriber("PRESSURE", "1", 950, 1030)
}
