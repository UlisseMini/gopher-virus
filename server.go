package main

import (
	"fmt"
	"os"
)

func main() {

}

func handle(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
