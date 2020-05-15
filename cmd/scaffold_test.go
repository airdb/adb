package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaffold(t *testing.T) {
	// 1tempDir, err := ioutil.TempDir(filepath.Join(Gopath, "src"), "test")
	tempDir, err := ioutil.TempDir(filepath.Join("/tmp", "src"), "test")

	if err != nil {
		t.Error("xxx")
		return
	}

	if !filepath.IsAbs(tempDir) {
		tempDir, err = filepath.Abs(tempDir)
		assert.NoError(t, err)
	}

	fmt.Printf("tempDir:%s\n", tempDir)
	assert.NoError(t, New(true).Generate(tempDir))

	defer os.RemoveAll(tempDir) // clean up
}
