package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
	"github.com/deexth/mvx/internal/ops"
)

func TestMoveFileIntoExistinDir(t *testing.T) {
	tmp := t.TempDir()

	srcFile := filepath.Join(tmp, "a.txt")
	err := os.WriteFile(srcFile, []byte("This is a test file, hello!"), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	dstDir := filepath.Join(tmp, "dir")
	err = os.Mkdir(dstDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	fS := fs.OSFS{}

	cfg := config.Config{
		HomeDir: tmp,
	}

	srcs, err := ops.HandlerSource([]string{srcFile}, fS)
	if err != nil {
		t.Fatal(err)
	}

	dst, err := ops.HandlerDestination(dstDir, cfg.HomeDir, fS)
	if err != nil {
		t.Fatal(err)
	}

	finalDst := ops.ResolveDestination(srcs[0], dst)

	err = fS.Rename(srcs[0].FullPath, finalDst)
	if err != nil {
		t.Fatal(err)
	}

	// original file should NOT exist
	if _, err := os.Stat(srcFile); !os.IsNotExist(err) {
		t.Fatalf("expected source to be moved, but it still exists")
	}

	// new file should exist
	expectedPath := filepath.Join(dstDir, "a.txt")
	if _, err := os.Stat(expectedPath); err != nil {
		t.Fatalf("expected file at destination, got error: %v", err)
	}
}
