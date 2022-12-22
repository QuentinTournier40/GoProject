package main

import "goproject/cmd/Captors"

func main() {
	Captors.RunCaptor("Pressure", "PRESSURE", "1", 950, 1030)
}
