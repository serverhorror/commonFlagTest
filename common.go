package commonFlagTest

import (
	"flag"
	"fmt"
	"os"
"runtime"
)

var (
	version string

	versionFlag = flag.Bool("version", false, "Print version and exit")
)

func Main() {
	flag.Parse()

	if *versionFlag {
		PrintVersion()
	}
}

func PrintVersion() {
	fmt.Printf("Version: %#q\n", version)
	fmt.Printf("Runtime: %#q\n", runtime.Version())
	os.Exit(1)
}
