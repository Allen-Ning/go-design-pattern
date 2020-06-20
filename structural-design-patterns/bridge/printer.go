package bridge

import "fmt"

type printer interface {
	printFile()
}

type hp struct{}

func (printer *hp) printFile() {
	fmt.Println("print by hp")
}

type epson struct{}

func (printer *epson) printFile() {
	fmt.Println("print by epson")
}
