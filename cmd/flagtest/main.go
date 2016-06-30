package main

import (
	"flag"
	"fmt"

	"github.com/serverhorror/commonFlagTest"
)

var (
	localFlag = flag.Bool("local", false, "Flag local to the main.go file")
)

func main() {
	commonFlagTest.Main()

	if *localFlag {
		fmt.Println("Local Flag")
	}
}
