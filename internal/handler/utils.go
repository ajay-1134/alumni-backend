package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
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

const (
	bucketName = "aura-poc-bucket"
)

func getImageUrl(c *gin.Context) (string, error) {
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
	wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	wc.ContentType = file.Header.Get("Content-Type")

	if _, err := io.Copy(wc, src); err != nil {
		log.Printf("failed to upload to GCS")
		return "", err
	}
	if err := wc.Close(); err != nil {
		log.Printf("failed to upload GCS writer")
		return "", err
	}

	signedURL, err := generateSignedURL(objectName)
	if err != nil {
		log.Printf("error occured in generating signed url")
		return "", err
	}

	return signedURL, nil
}

func generateSignedURL(objectName string) (string, error) {
	serviceAccount := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	googleAccessID,err := extractClientEmail(serviceAccount)
	if err != nil {
		return "",err
	}

	privateKey,err := extractPrivateKey(serviceAccount)
	if err != nil {
		return "",err
	}

	url, err := storage.SignedURL(bucketName, objectName, &storage.SignedURLOptions{
		Method:         "GET",
		Expires:        time.Now().Add(1 * time.Hour),
		GoogleAccessID: googleAccessID,
		PrivateKey:     privateKey,
	})

	return url, err
}

type ServiceAccount struct {
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
}

func extractClientEmail(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("error occured in reading credential file")
		return "", err
	}

	var sa ServiceAccount

	err = json.Unmarshal(data, &sa)
	if err != nil {
		log.Printf("error occured in unmarshalling json file")
		return "", err
	}

	return sa.ClientEmail, nil
}

func extractPrivateKey(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("error occured in reading credential file")
		return nil, err
	}

	var sa ServiceAccount
	err = json.Unmarshal(data, &sa)
	if err != nil {
		log.Printf("error occured in unmarshalling json file")
		return nil, err
	}

	return []byte(sa.PrivateKey), nil
}
