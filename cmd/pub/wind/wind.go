package main

import (
	"goproject/internal/publishers"
)

func main() {
	publishers.RunPublisher("WIND", 2, 0, 250)
}
