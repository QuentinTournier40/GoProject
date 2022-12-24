package main

import (
	"goproject/cmd/pub/Subscribers"
)

func main() {
	Subscribers.RunSubscriber("WIND", "3", 0, 250)
}
