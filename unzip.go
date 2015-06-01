// unzip is a package for extracting ZIP files
package unzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"errors"
)

// Given a source filename and a destination path, extract the ZIP archive
func Extract(zipFilename, destPath string) error {

	// Open the source filename for reading
	zipReader, err := zip.OpenReader(zipFilename)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// For each file in the archive
	for _, archiveReader := range zipReader.File {

		// Open the file in the archive
		archiveFile, err := archiveReader.Open()
		if err != nil {
			return err
		}
		defer archiveFile.Close()

		// Prepare to write the file
		finalPath := filepath.Join(destPath, archiveReader.Name)

		// Check if the file to extract is just a directory
		if archiveReader.FileInfo().IsDir() {
			err = os.MkdirAll(finalPath, 0755)
			if err != nil {
				return err
			}
			continue
		}
		if len(archiveReader.Name) >= 150 {
			return errors.New("Too long filename: " + archiveReader.Name)
		}

		// Create all needed directories
		if os.MkdirAll(filepath.Dir(finalPath), 0755) != nil {
			return err
		}

		// Prepare to write the destination file
		destinationFile, err := os.OpenFile(finalPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, archiveReader.Mode())
		if err != nil {
			return err
		}
		defer destinationFile.Close()

		// Write the destination file
		if _, err = io.Copy(destinationFile, archiveFile); err != nil {
			return err
		}
	}

	return nil
}
