package main

import "github.com/svcodestore/sv-sso-gin/core"

// go:generate go env -w GO111MODULE=on
// go:generate go env -w GOPROXY=https://goproxy.cn,direct
// go:generate go mod tidy
// go:generate go mod download
func main() {
	core.RunServer()
}
