package capture

import (
	"bytes"
	"runtime/pprof"
	"time"
)

// CaptureCPUProfile 采集CPU Profile数据
func CaptureCPUProfile(duration time.Duration) ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		return nil, err
	}
	time.Sleep(duration)
	pprof.StopCPUProfile()
	return buf.Bytes(), nil
}

// CaptureMemoryProfile 采集Memory Profile数据
func CaptureMemoryProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.WriteHeapProfile(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureBlockProfile 采集Block Profile数据
func CaptureBlockProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("block").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureMutexProfile 采集Mutex Profile数据
func CaptureMutexProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("mutex").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureGoroutineProfile 采集Goroutine Profile数据
func CaptureGoroutineProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("goroutine").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
