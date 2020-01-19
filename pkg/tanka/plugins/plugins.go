package plugins

import (
	"fmt"
	"plugin"

	"github.com/google/go-jsonnet"
)

const PluginSymbol = "TankaExtensions"

// ImportPlugin extract Jsonnet native function from a Go dynamic
// library (plugin).
func ImportPlugin(path string) ([]*jsonnet.NativeFunction, error) {
	ext, err := plugin.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open plugin Tanka: %w", err)
	}

	nativeFuncs, err := ext.Lookup(PluginSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to load '%s' in '%s': %w", PluginSymbol, path, err)
	}

	if _, valid := nativeFuncs.(func() []*jsonnet.NativeFunction); !valid {
		return nil, fmt.Errorf("failed to use '%s': %s must be %T", path, PluginSymbol, (func() []*jsonnet.NativeFunction)(nil))
	}

	return nativeFuncs.(func() []*jsonnet.NativeFunction)(), nil
}
