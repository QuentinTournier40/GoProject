package main

import (
	"goproject/cmd/pub/Subscribers"
)

func main() {
	Subscribers.RunSubscriber("TEMPERATURE", "2", -15, 43)
}
