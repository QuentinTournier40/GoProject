package main

import (
	"goproject/cmd/pub/Publishers"
)

func main() {
	Publishers.RunPublisher("PRESSURE", "1", 950, 1030)
}
