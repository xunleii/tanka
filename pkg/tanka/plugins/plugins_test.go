package plugins

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestImportPlugin(t *testing.T) {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "dummy.tkplugin.so", "./testdata")

	var errBuf, outBuff bytes.Buffer
	cmd.Stderr = &errBuf
	cmd.Stdout = &outBuff
	err := cmd.Run()
	require.NoError(t, err)

	defer func() {
		_ = os.Remove("dummy.tkplugin.so")
	}()

	funcs, err := ImportPlugin("dummy.tkplugin.so")
	require.NoError(t, err)
	require.Len(t, funcs, 1)
	assert.Equal(t, funcs[0].Name, "helloWorld")
}
