versiontest:
	# WRONG! go run -a -ldflags="-X=commonFlagTest.Version='my-uber-version'" cmd/flagtest/main.go -version
	go run -ldflags="-X github.com/serverhorror/commonFlagTest.versionString='$(shell TZ=UTC date '+%FT%T.%N%:z')' -X main.localVersion='boo'" cmd/flagtest/main.go -local-version
	go run -ldflags="-X github.com/serverhorror/commonFlagTest.versionString='$(shell TZ=UTC date '+%FT%T.%N%:z')' -X main.localVersion='boo'" cmd/flagtest/main.go -version
