package main

import "base"

func main() {
	err := base.PrintExactTime()
	if err != nil {
		return
	}
}
