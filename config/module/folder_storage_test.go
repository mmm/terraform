package module

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFolderStorage_impl(t *testing.T) {
	var _ Storage = new(FolderStorage)
}

func TestFolderStorage(t *testing.T) {
	s := &FolderStorage{StorageDir: tempDir(t)}

	module := testModule("basic")

	// A module shouldn't exist at first...
	_, ok, err := s.Dir(module)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if ok {
		t.Fatal("should not exist")
	}

	// We can get it
	err = s.Get(module, false)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	// Now the module exists
	dir, ok, err := s.Dir(module)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if !ok {
		t.Fatal("should exist")
	}

	mainPath := filepath.Join(dir, "main.tf")
	if _, err := os.Stat(mainPath); err != nil {
		t.Fatalf("err: %s", err)
	}

}

func tempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "tf")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if err := os.RemoveAll(dir); err != nil {
		t.Fatalf("err: %s", err)
	}

	return dir
}
