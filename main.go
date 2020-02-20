package main

import (
	"os"

	"github.com/jameshiew/same/cmd"
)

const logLevelEnvKey = "SAME_LOG_LEVEL"

var (
	version  string
	logLevel string
)

func init() {
	if version == "" {
		version = "dev"
	}
	logLevel = os.Getenv(logLevelEnvKey)
	if logLevel == "" {
		logLevel = "WARN"
	}
}

func main() {
	cmd.Execute(version, logLevel)
}
