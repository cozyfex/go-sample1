package sample1

import (
	"fmt"

	"cozyfex.com/sample1/sub"
	"cozyfex.com/sample1/sub/directory"
)

func Name() {
	sub.Who()
	directory.Where()
	fmt.Println("This is sample1 module")
}

func What() {
	fmt.Println("What")
}
