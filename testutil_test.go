package main

import (
	"os"
	"path/filepath"
	"testing"
)

// writeHullImage is the shared unit-test helper that dumps the current hull
// image state to testdata/out/<name>.png (created as needed).
func writeHullImage(t *testing.T, name string, h *hull) {
	t.Helper()
	dir := filepath.Join("testdata", "out")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", dir, err)
	}
	path := filepath.Join(dir, name+".png")
	if err := WriteHullImage(path, h); err != nil {
		t.Fatalf("WriteHullImage(%s): %v", path, err)
	}
	t.Logf("wrote visualization: %s", path)
}
