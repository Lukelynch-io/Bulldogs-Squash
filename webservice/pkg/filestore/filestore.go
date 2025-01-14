package filestore

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

var filestoreLocation string

func SetFilestoreLocation(path string) {
	filestoreLocation = path
}

func extractFilenameExtension(filename string) string {
	extensionStart := -1
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			extensionStart = i
			break
		}
	}

	if extensionStart == -1 {
		return ""
	}
	return filename[extensionStart:]
}

func StoreFile(file *multipart.FileHeader) (string, error) {
	newFileName := uuid.NewString()
	completeFilePath := filestoreLocation + newFileName + extractFilenameExtension(file.Filename)
	// TODO: validate filepath
	inputfile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer inputfile.Close()
	createdFile, err := os.Create(completeFilePath)
	if err != nil {
		return "", err
	}
	defer createdFile.Close()

	if _, err := io.Copy(createdFile, inputfile); err != nil {
		return "", err
	}

	return completeFilePath, nil
}
