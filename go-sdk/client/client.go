package client

import "time"

type PProfClient struct {
	uploader *Uploader
}

func NewPProfClient(uploadURL string) *PProfClient {
	return &PProfClient{
		uploader: NewUploader(uploadURL),
	}
}

func (c *PProfClient) CollectCPUProfileAndUpload(duration time.Duration) error {
	data, err := CaptureCPUProfile(duration)
	if err != nil {
		return err
	}
	return c.uploader.Upload(data)
}

func (c *PProfClient) CollectMemoryProfileAndUpload() error {
	data, err := CaptureMemoryProfile()
	if err != nil {
		return err
	}
	return c.uploader.Upload(data)
}

func (c *PProfClient) CollectBlockProfileAndUpload() error {
	data, err := CaptureBlockProfile()
	if err != nil {
		return err
	}
	return c.uploader.Upload(data)
}

func (c *PProfClient) CollectMutexProfileAndUpload() error {
	data, err := CaptureMutexProfile()
	if err != nil {
		return err
	}
	return c.uploader.Upload(data)
}

func (c *PProfClient) CollectGoroutineProfileAndUpload() error {
	data, err := CaptureGoroutineProfile()
	if err != nil {
		return err
	}
	return c.uploader.Upload(data)
}
