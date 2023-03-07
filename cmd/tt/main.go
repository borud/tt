package main

import (
	"time"

	"github.com/borud/tt/pkg/util"
)

var opt struct {
	GRPCAddr    string        `long:"grpc-addr" default:":6661" description:"gRPC endpoint addr" required:"yes"`
	HTTPAddr    string        `long:"http-addr" default:":6660" description:"HTTP endpoint addr" required:"yes"`
	GRPCTimeout time.Duration `long:"timeout" default:"10s" description:"gRPC timeout"`

	Server  serverCmd  `command:"server" description:"start the server"`
	User    userCmd    `command:"user" description:"user management commands"`
	Project projectCmd `command:"project" description:"project management commands"`
	Work    workCmd    `command:"work" description:"work entry commands"`
	Snippet snippetCmd `command:"snippet" alias:"snip" description:"snippet entry commands"`
}

func main() {
	util.FlagParse(&opt)
}
