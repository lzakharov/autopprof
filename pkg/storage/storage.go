// Package storage provides a file storage.
package storage

import (
	"io"
)

// Storage represents a file storage.
type Storage interface {
	Open(name string) (File, error)
}

// File represents a writable file.
type File interface {
	io.Writer
	io.Closer
}
