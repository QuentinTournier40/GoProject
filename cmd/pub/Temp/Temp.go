package main

import (
	"goproject/cmd/pub/Captors"
)

func main() {
	Captors.RunCaptor("Temp", "TEMPERATURE", "2", -15, 43)
}
