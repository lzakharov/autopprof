package storage

import "os"

// FS is a file system storage.
type FS struct{}

// Open creates or truncates the named file.
func (f FS) Open(name string) (File, error) {
	return os.Create(name)
}
