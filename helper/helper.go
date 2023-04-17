package helper

import (
	"errors"
	"fmt"
	"log"
	"main/fn"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	Dir = "DIR"
)

func SaveFile(c *gin.Context, file *multipart.FileHeader, err error) (string, error) {
	if err != nil {
		return "", err
	}

	fileExt := filepath.Ext(file.Filename)

	// Check the file type
	mimeType := mime.TypeByExtension(fileExt)
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return "", errors.New("invalid image file type")
	}

	// Generate a unique filename
	fileId := fn.GenerateUUID()
	fileName := fmt.Sprintf("%s%s", fileId, strings.ToLower(fileExt))

	// Save the file to the server
	dst := fmt.Sprintf("%s/%s/%s", os.Getenv(Dir), "upload", fileName)
	saveErr := c.SaveUploadedFile(file, dst)
	if saveErr != nil {
		return "", saveErr
	}

	return fileName, nil
}

func RemoveFile(fileName string) (bool, error) {
	if fileName == "" {
		return false, nil
	}

	rm := fmt.Sprintf("%s/%s/%s", os.Getenv(Dir), "upload", fileName)
	if err := os.Remove(rm); err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}
