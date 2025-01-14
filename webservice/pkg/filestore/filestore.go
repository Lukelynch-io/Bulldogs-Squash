package filestore

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type Filestore interface {
	StoreFile(file *multipart.FileHeader) (string, error)
}

type LocalFilestore struct {
	filestoreLocation string
}

func (fs *LocalFilestore) SetFilestoreLocation(path string) {
	fs.filestoreLocation = path
}

func (fs *LocalFilestore) StoreFile(file *multipart.FileHeader) (string, error) {
	newFileName := uuid.NewString()
	completeFilePath := fs.filestoreLocation + newFileName + extractFilenameExtension(file.Filename)
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
