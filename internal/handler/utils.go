package handler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ajay-1134/alumni-backend/internal/constants"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type DateOnly struct {
	time.Time
}

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // remove quotes
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return errors.New("invalid date format, use YYYY-MM-DD")
	}
	d.Time = t
	return nil
}

// Convert back to JSON string
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}

func stringToUint(s string) (uint, error) {
	val64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val64), nil
}

func getUserID(c *gin.Context) (uint, error) {
	id := c.Param("id")

	if id != "" {
		id := c.Param("id")
		return stringToUint(id)
	}

	userId, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("userID key does not exist")
	}

	return userId.(uint), nil
}

func uploadImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		log.Printf("image file is not provided")
		return "", nil
	}

	src, err := file.Open()
	if err != nil {
		log.Printf("error occured in opening the file")
		return "", err
	}
	defer src.Close()

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")))
	if err != nil {
		log.Printf("error occured in initializing gcs client")
		return "", err
	}
	defer client.Close()

	objectName := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), file.Filename)
	wc := client.Bucket(constants.BucketName).Object(objectName).NewWriter(ctx)
	wc.ContentType = file.Header.Get("Content-Type")

	if _, err := io.Copy(wc, src); err != nil {
		log.Printf("failed to upload to GCS")
		return "", err
	}
	if err := wc.Close(); err != nil {
		log.Printf("failed to upload GCS writer")
		return "", err
	}

	return objectName, nil
}
