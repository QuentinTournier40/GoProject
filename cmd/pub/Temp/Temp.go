package main

import (
	"goproject/cmd/pub/Publishers"
)

func main() {
	Publishers.RunPublisher("TEMPERATURE", 1, -15, 43)
}
