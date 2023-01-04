package main

import (
	"goproject/internal/Subscribers"
)

func main() {
	Subscribers.RunSubscriber("client-csv", false)
}
