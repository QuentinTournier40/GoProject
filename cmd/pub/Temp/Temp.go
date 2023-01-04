package main

import (
	"goproject/internal/Publishers"
)

func main() {
	Publishers.RunPublisher("TEMPERATURE", 1, -15, 43)
}
