package main

import "fmt"

const version = "0.0.0"

var revision = ""

func FmtVersion() string {
	if revision == "" {
		return version
	}
	return fmt.Sprintf("%s, build %s", version, revision)
}
