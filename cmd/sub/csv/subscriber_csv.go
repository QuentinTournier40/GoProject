package main

import (
	"goproject/cmd/sub/Subscribers"
)

func main() {
	Subscribers.RunSubscriber("client-csv", false)
}
