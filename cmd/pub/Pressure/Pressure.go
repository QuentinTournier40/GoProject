package main

import (
	"goproject/internal/Publishers"
)

func main() {
	Publishers.RunPublisher("PRESSURE", 0, 950, 1030)
}
