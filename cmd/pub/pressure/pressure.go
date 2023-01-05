package main

import (
	"goproject/internal/publishers"
)

func main() {
	publishers.RunPublisher("PRESSURE", 0, 950, 1030)
}
