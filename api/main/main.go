package main

import "goproject/api/router"

func main() {
	a := router.App{}
	a.Initialize()
	a.Run(":8080")
}
