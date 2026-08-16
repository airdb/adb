package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteBashCompletionScriptIncludesCompatShim(t *testing.T) {
	var buf bytes.Buffer

	if err := writeBashCompletionScript(&buf); err != nil {
		t.Fatalf("writeBashCompletionScript returned error: %v", err)
	}

	script := buf.String()
	if !strings.Contains(script, "if ! declare -F _get_comp_words_by_ref >/dev/null 2>&1; then") {
		t.Fatal("expected completion script to include _get_comp_words_by_ref compatibility shim")
	}
	if !strings.Contains(script, "__adb_init_completion()") {
		t.Fatal("expected completion script to include cobra-generated bash completion")
	}
}
