package cmd

import (
	"bytes"
	"encoding/base64"
	"strings"
	"testing"
)

func TestJWTCommand(t *testing.T) {
	token := "header." + base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"123","admin":true}`)) + ".signature"
	for _, input := range []string{token, "  Bearer " + token + "  "} {
		var output bytes.Buffer
		rootCmd.SetOut(&output)
		rootCmd.SetErr(&output)
		rootCmd.SetArgs([]string{"jwt", input})
		err := rootCmd.Execute()
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := output.String(), "{\n  \"admin\": true,\n  \"sub\": \"123\"\n}\n"; got != want {
			t.Fatalf("output = %q, want %q", got, want)
		}
	}
}

func TestJWTCommandRejectsInvalidInput(t *testing.T) {
	for _, tc := range []struct {
		args []string
		want string
	}{
		{[]string{"a", "b"}, "accepts at most 1 arg(s)"},
		{[]string{"a.b"}, "invalid JWT: expected 3 parts"},
		{[]string{"a.!.c"}, "decode payload:"},
		{[]string{"a.bm90LWpzb24.c"}, "parse JSON:"},
	} {
		var output bytes.Buffer
		rootCmd.SetOut(&output)
		rootCmd.SetErr(&output)
		rootCmd.SetArgs(append([]string{"jwt"}, tc.args...))
		err := rootCmd.Execute()
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
		if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Fatalf("args %v: error = %v, want %q", tc.args, err, tc.want)
		}
	}
}

func TestJWTCommandWithoutArgsShowsUsage(t *testing.T) {
	var output bytes.Buffer
	rootCmd.SetOut(&output)
	rootCmd.SetErr(&output)
	rootCmd.SetArgs([]string{"jwt"})
	defer func() {
		rootCmd.SetOut(nil)
		rootCmd.SetErr(nil)
		rootCmd.SetArgs(nil)
	}()
	if err := rootCmd.Execute(); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"Usage:",
		"adb jwt <jwt-token>",
		`adb jwt "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTYiLCJuYW1lIjoiRGVhbiIsImFkbWluIjp0cnVlfQ.signature"`,
	} {
		if !strings.Contains(output.String(), want) {
			t.Errorf("help missing %q", want)
		}
	}
}
