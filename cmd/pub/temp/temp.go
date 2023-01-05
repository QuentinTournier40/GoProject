package main

import (
	"goproject/internal/publishers"
)

func main() {
	publishers.RunPublisher("TEMPERATURE", 1, -15, 43)
}
