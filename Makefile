versiontest:
	go run -ldflags="-X commonFlagTest.version='my-uber-version'" cmd/flagtest/main.go -version
