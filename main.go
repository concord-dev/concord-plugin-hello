// Reference Concord plugin for protocol v1.
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	plugin "github.com/concord-dev/concord-plugin-sdk/plugin"
)

type hello struct{}

func (hello) Capabilities() plugin.Capabilities {
	return plugin.Capabilities{
		Source:         "hello",
		Version:        "v0.1.0",
		SupportedTypes: []string{"greeting"},
		DocsURL:        "https://github.com/concord-dev/concord-plugin-hello",
	}
}

func (hello) Probe(_ context.Context) (string, error) {
	return fmt.Sprintf("hello plugin OK (go %s, %s/%s)", runtime.Version(), runtime.GOOS, runtime.GOARCH), nil
}

func (hello) Collect(_ context.Context, ref plugin.EvidenceRef) (any, error) {
	if ref.Type != "greeting" {
		return nil, plugin.ErrUnsupportedType
	}
	name := plugin.StringParam(ref, "name")
	if name == "" {
		name = "world"
	}
	return map[string]any{
		"message":        fmt.Sprintf("hello, %s", name),
		"fetched_at":     time.Now().UTC().Format(time.RFC3339),
		"protocol":       plugin.ProtocolVersion,
		"plugin_name":    "hello",
		"plugin_version": "v0.1.0",
	}, nil
}

func main() { plugin.Serve(hello{}) }
