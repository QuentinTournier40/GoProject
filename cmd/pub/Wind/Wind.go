package main

import (
	"goproject/cmd/pub/Publishers"
)

func main() {
	Publishers.RunPublisher("WIND", 2, 0, 250)
}
