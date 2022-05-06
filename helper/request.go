package helper

import (
	"errors"
	uuid2 "github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var mime = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
}

func UploadImage(r *http.Request, maxSize int, path string) (string, error) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		return "", err
	}

	uploadedFile, header, err := r.FormFile("image")
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	if header.Size > int64(maxSize)*1000*1000 {
		return "", errors.New("file to large")
	}

	ext, ok := mime[header.Header["Content-Type"][0]]
	if !ok {
		return "", errors.New("mime type not supported")
	}

	uuid, err := uuid2.NewUUID()
	if err != nil {
		return "", err
	}
	filename := uuid.String() + "." + ext

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fileLocation := filepath.Join(dir, path, filename)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, uploadedFile)
	if err != nil {
		return "", err
	}

	return filename, nil
}
