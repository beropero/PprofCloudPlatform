package client

import (
	"bytes"
	"errors"
	"net/http"
	"time"
)

type Uploader struct {
	uploadURL string
	client    *http.Client
}

func NewUploader(uploadURL string) *Uploader {
	return &Uploader{
		uploadURL: uploadURL,
		client:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (u *Uploader) Upload(data []byte) error {
	req, err := http.NewRequest("POST", u.uploadURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := u.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("upload failed")
	}

	return nil
}
