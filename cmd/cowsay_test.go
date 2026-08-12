package cmd

import (
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"testing"
)

func TestLoadCustomCowsayFiguresMissingDir(t *testing.T) {
	tempHome := t.TempDir()
	previousHomeDir := cowsayUserHomeDir
	cowsayUserHomeDir = func() (string, error) {
		return tempHome, nil
	}
	t.Cleanup(func() {
		cowsayUserHomeDir = previousHomeDir
	})

	figures, err := loadCustomCowsayFigures()
	if err != nil {
		t.Fatalf("loadCustomCowsayFigures returned error: %v", err)
	}
	if len(figures) != 0 {
		t.Fatalf("expected no figures, got %d", len(figures))
	}
}

func TestLoadCustomCowsayFiguresRecursive(t *testing.T) {
	tempHome := t.TempDir()
	previousHomeDir := cowsayUserHomeDir
	cowsayUserHomeDir = func() (string, error) {
		return tempHome, nil
	}
	t.Cleanup(func() {
		cowsayUserHomeDir = previousHomeDir
	})

	customDir := filepath.Join(tempHome, cowsayCustomFigureDir, "animals")
	if err := os.MkdirAll(customDir, 0o755); err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(customDir, "yak.txt"), []byte(" \\\\\n(oo)"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	figures, err := loadCustomCowsayFigures()
	if err != nil {
		t.Fatalf("loadCustomCowsayFigures returned error: %v", err)
	}

	got, ok := figures["animals/yak.txt"]
	if !ok {
		t.Fatalf("expected figure animals/yak.txt, got %v", cowsayFigureNames(figures))
	}

	wantBody := []string{" \\\\", "(oo)"}
	if !reflect.DeepEqual(got.body, wantBody) {
		t.Fatalf("unexpected figure body: got %#v want %#v", got.body, wantBody)
	}
}

func TestLoadCowsayFiguresCustomOverridesBuiltin(t *testing.T) {
	tempHome := t.TempDir()
	previousHomeDir := cowsayUserHomeDir
	cowsayUserHomeDir = func() (string, error) {
		return tempHome, nil
	}
	t.Cleanup(func() {
		cowsayUserHomeDir = previousHomeDir
	})

	customDir := filepath.Join(tempHome, cowsayCustomFigureDir)
	if err := os.MkdirAll(customDir, 0o755); err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}
	if err := os.WriteFile(filepath.Join(customDir, "cow"), []byte("custom cow"), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	figures, err := loadCowsayFigures()
	if err != nil {
		t.Fatalf("loadCowsayFigures returned error: %v", err)
	}

	got, ok := figures["cow"]
	if !ok {
		t.Fatal("expected builtin cow figure to exist")
	}

	wantBody := []string{"custom cow"}
	if !reflect.DeepEqual(got.body, wantBody) {
		t.Fatalf("expected custom cow to override builtin: got %#v want %#v", got.body, wantBody)
	}
}

func TestCowsayFortunesIncludeChineseLines(t *testing.T) {
	want := []string{
		"行而不辍，未来可期。",
		"路虽远，行则将至。",
		"守正而后出新。",
		"见微知著，行稳致远。",
	}

	for _, line := range want {
		if !slices.Contains(cowsayFortunes, line) {
			t.Fatalf("expected cowsayFortunes to include %q", line)
		}
	}
}
