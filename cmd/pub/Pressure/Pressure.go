package main

import (
	"goproject/cmd/pub/Publishers"
)

func main() {
	Publishers.RunPublisher("PRESSURE", 0, 950, 1030)
}
