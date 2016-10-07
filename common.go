package commonFlagTest

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

var versionString string

var (
	versionFlag = flag.Bool("version", false, "Print version and exit")
)

func init() {
	if versionString == "" {
		versionString = "<development>"
	}
}

func Main() {
	flag.Parse()

	if *versionFlag {
		PrintVersion()
	}
}

func PrintVersion() {
	fmt.Printf("Version: %#q\n", versionString)
	fmt.Printf("Runtime: %#q\n", runtime.Version())
	os.Exit(0)
}
