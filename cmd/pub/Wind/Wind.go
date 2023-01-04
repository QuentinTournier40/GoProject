package main

import (
	"goproject/internal/Publishers"
)

func main() {
	Publishers.RunPublisher("WIND", 2, 0, 250)
}
