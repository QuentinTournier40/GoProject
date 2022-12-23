package main

import (
	"goproject/cmd/pub/Captors"
)

func main() {
	Captors.RunCaptor("WIND", "3", 0, 250)
}
