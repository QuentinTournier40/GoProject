package main

import (
	"goproject/internal/subscribers"
)

func main() {
	subscribers.RunSubscriber("client-api", true)
}
