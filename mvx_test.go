package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/deexth/mvx/internal/cli"
	"github.com/deexth/mvx/internal/config"
	"github.com/deexth/mvx/internal/fs"
	"github.com/deexth/mvx/internal/ops"
)

func TestMVXBasicMVFunctionality(t *testing.T) {
	cases := []struct {
		name        string
		setup       func(root string) (src []string, dst string)
		expectedErr bool
		validate    func(t *testing.T, root string)
	}{
		{
			name: "renaming file",
			setup: func(root string) (src []string, dst string) {
				os.WriteFile(filepath.Join(root, "a.txt"), []byte("This is a test file!"), 0o644)
				return []string{"a.txt"}, "baz.txt"
			},
			expectedErr: false,
			validate: func(t *testing.T, root string) {
				if _, err := os.Stat(filepath.Join(root, "baz.txt")); err != nil {
					t.Fatalf("file not renamed: %v", err)
				}
			},
		},
		{
			name: "moving file into existing directory",
			setup: func(root string) (src []string, dst string) {
				os.WriteFile(filepath.Join(root, "a.txt"), []byte("This is a test file!"), 0o644)
				os.Mkdir(filepath.Join(root, "foo"), 0o755)
				return []string{"a.txt"}, "foo"
			},
			expectedErr: false,
			validate: func(t *testing.T, root string) {
				if _, err := os.Stat(filepath.Join(root, "foo", "a.txt")); err != nil {
					t.Fatalf("file not moved: %v", err)
				}
			},
		},
		{
			name: "create nested destination",
			setup: func(root string) (src []string, dst string) {
				os.WriteFile(filepath.Join(root, "a.txt"), []byte("This is a test file!"), 0o644)
				return []string{"a.txt"}, "foo/bar/baz.txt"
			},
			expectedErr: true,
		},
		{
			name: "mutliple sources to a file destination",
			setup: func(root string) (src []string, dst string) {
				os.WriteFile(filepath.Join(root, "a.txt"), []byte("hi"), 0o644)
				os.WriteFile(filepath.Join(root, "b.txt"), []byte("hi"), 0o644)
				return []string{"a.txt", "b.txt"}, "target.txt"
			},
			expectedErr: true,
		},
		{
			name: "tilde as destination or prefix",
			setup: func(root string) (src []string, dst string) {
				os.WriteFile(filepath.Join(root, "a.txt"), []byte("This is a test file!"), 0o644)
				return []string{"a.txt"}, "~"
			},
			expectedErr: false,
			validate: func(t *testing.T, root string) {
				if _, err := os.Stat(filepath.Join(root, "a.txt")); err != nil {
					t.Fatalf("file not moved: %v", err)
				}
			},
		},
		{
			name: "tilde as prefix",
			setup: func(root string) (src []string, dst string) {
				os.WriteFile(filepath.Join(root, "a.txt"), []byte("This is a test file!"), 0o644)
				return []string{"a.txt"}, "~/baz.txt"
			},
			expectedErr: false,
			validate: func(t *testing.T, root string) {
				if _, err := os.Stat(filepath.Join(root, "baz.txt")); err != nil {
					t.Fatalf("file not moved: %v", err)
				}
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			cwd, _ := os.Getwd()

			// Overide HOME
			t.Setenv("HOME", root)

			os.Chdir(root)
			defer os.Chdir(cwd)

			fS := fs.OSFS{}
			cfg := config.Config{
				CWD: root,
			}
			src, dst := tt.setup(root)
			cfg.Source = src
			cfg.Destination = dst

			err := ops.Move(&cfg, cli.MoveOptions{}, fS)

			if tt.expectedErr && err == nil {
				t.Fatal("expected an err and got nil")
			}

			if !tt.expectedErr && err != nil {
				t.Fatal(err)
			}

			if tt.validate != nil {
				tt.validate(t, root)
			}
		})

	}

}
