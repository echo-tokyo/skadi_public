// Package file (utils) contains help-functions for file uploading.
package file

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
)

const _mpfdFileKey = "file" // key for mpfd files

// ParseAndSaveFiles parses files from mpfd, saves the to file system
func ParseAndSaveFiles(ctx *fiber.Ctx, fileDir string) (entity.Files, error) {
	// parse mpfd
	mpfd, err := ctx.MultipartForm()
	if err != nil {
		return nil, fmt.Errorf("parse mpfd: %w", err)
	}
	if mpfd == nil {
		return nil, nil
	}

	uploadedFiles := make(entity.Files, len(mpfd.File[_mpfdFileKey]))
	for idx, rawFile := range mpfd.File[_mpfdFileKey] {
		// collect file's metadata
		uploadedFiles[idx] = entity.NewFile(
			rawFile.Filename,
			rawFile.Header["Content-Type"][0],
			rawFile.Size,
			entity.FileWithPathPrefix(fileDir),
		)
		// save file to file system
		if err := ctx.SaveFile(rawFile, uploadedFiles[idx].Path); err != nil {
			uploadedFiles.Cleanup()
			return nil, fmt.Errorf("save file %s: %w", rawFile.Filename, err)
		}
	}
	return uploadedFiles, nil
}
