package main

import "fmt"

var (
	build   string
	version string
)

func init() {
	if build == "" {
		build = "unknown"
	}
	if version == "" {
		version = "0.0.0-dev"
	}
}

func main() {
	fmt.Println(fmt.Sprintf("Hello, GopherCon Turkey! This is build %s and version %s", build, version))
}
