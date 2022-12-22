package main

import (
	"goproject/cmd/pub/Captors"
)

func main() {
	Captors.RunCaptor("Wind", "WIND", "3", 0, 250)
}
