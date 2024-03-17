package main

import (
	"bookstack/cmd"
)

var (
	version  = "UNKNOWN"
	revision = "UNKNOWN"
)

func main() {
	cmd.Version = version
	cmd.Revision = revision
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
