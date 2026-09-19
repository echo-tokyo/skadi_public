package entity

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// File represents a metadata for the file.
type File struct {
	// file ID
	ID int `json:"id" validate:"required"`
	// filename
	Name string `json:"name" validate:"required"`
	// file MIME-type
	MimeType string `json:"mime_type" validate:"required"`
	// file size in bytes
	Size int64 `json:"size" validate:"required"`
	// filepath
	Path string `json:"-"`
	// prefix for filepath
	pathPrefix string `json:"-"`
}

// TableName determines DB table name for the file object.
func (*File) TableName() string {
	return "file"
}

type FileOption func(*File)

// NewFile returns a new instance of File.
func NewFile(name, mimeType string, size int64, options ...FileOption) *File {
	file := &File{
		Name:     name,
		MimeType: mimeType,
		Size:     size,
	}
	// apply options
	for _, option := range options {
		option(file)
	}
	// generate a new filename
	file.Path = filepath.Join(file.pathPrefix, uuid.NewString())
	return file
}

// FileWithPathPrefix sets custom prefix for file path.
func FileWithPathPrefix(s string) FileOption {
	return func(f *File) {
		f.pathPrefix = s
	}
}

// Remove removes the saved file from the file system.
func (f *File) Remove() {
	if err := os.Remove(f.Path); err != nil {
		slog.Warn("remove file %s: %w", f.Path, err)
	}
	slog.Info("remove file",
		"filename", f.Name,
		"mime-type", f.MimeType,
		"size", f.Size,
		"path", f.Path)
}

// Files represents a slice of File objects.
type Files []*File

// Cleanup removes all files.
func (f Files) Cleanup() {
	for idx := range f {
		f[idx].Remove()
	}
}
