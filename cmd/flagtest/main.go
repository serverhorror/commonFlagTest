package main

import (
	"flag"
	"fmt"

	"github.com/serverhorror/commonFlagTest"
)

var localVersion string

var (
	localFlag        = flag.Bool("local", false, "Flag local to the main.go file")
	localVersionFlag = flag.Bool("local-version", false, "Flag to use a locally injected version")
)

func init() {
	if localVersion == "" {
		localVersion = "<development>"
	}
}

func main() {
	commonFlagTest.Main()

	if *localFlag {
		fmt.Println("Local Flag")
	}
	if *localVersionFlag {
		fmt.Println("Local Version Flag")
		fmt.Printf("Local Version: %#q\n", localVersion)
	}
}
