package main

import (
	"goproject/cmd/pub/Captors"
)

func main() {
	Captors.RunCaptor("PRESSURE", "1", 950, 1030)
}
