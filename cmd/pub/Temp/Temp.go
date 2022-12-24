package main

import (
	"goproject/cmd/pub/Publishers"
)

func main() {
	Publishers.RunPublisher("TEMPERATURE", "2", -15, 43)
}
