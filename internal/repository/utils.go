package repository

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ajay-1134/alumni-backend/internal/constants"
)

func generateSignedURL(objectName string) (string, error) {
	serviceAccount := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	googleAccessID, err := extractClientEmail(serviceAccount)
	if err != nil {
		return "", err
	}

	privateKey, err := extractPrivateKey(serviceAccount)
	if err != nil {
		return "", err
	}

	url, err := storage.SignedURL(constants.BucketName, objectName, &storage.SignedURLOptions{
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
