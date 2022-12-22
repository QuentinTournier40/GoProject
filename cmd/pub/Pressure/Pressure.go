package main

import (
	"goproject/cmd/pub/Captors"
)

func main() {
	Captors.RunCaptor("Pressure", "PRESSURE", "1", 950, 1030)
}
