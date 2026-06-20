// Command concord-plugin-hello is the reference Concord plugin for protocol v1.
package main

import (
	"github.com/concord-dev/concord-plugin-hello/internal/hello"
	plugin "github.com/concord-dev/concord-plugin-sdk/plugin"
)

func main() { plugin.Serve(hello.New()) }
