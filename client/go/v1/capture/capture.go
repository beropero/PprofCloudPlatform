package capture

import (
	"bytes"
	"runtime/pprof"
	"runtime/trace"
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

// CaptureThreadProfile 采集Thread Profile数据
func CaptureThreadProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("threadcreate").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureHeapProfile 采集Heap Profile数据
func CaptureHeapProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("heap").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureAllocsProfile 采集Allocs Profile数据
func CaptureAllocsProfile() ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.Lookup("allocs").WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureTraceProfile 采集Trace Profile数据
func CaptureTraceProfile(duration time.Duration) ([]byte, error) {
	var buf bytes.Buffer
	if err := trace.Start(&buf); err != nil {
		return nil, err
	}
	time.Sleep(duration)
	trace.Stop()
	return buf.Bytes(), nil
}

// CaptureCustomProfile 采集自定义Profile数据
func CaptureCustomProfile() ([]byte, error) {
	var customProfile = pprof.NewProfile("custom")
	var buf bytes.Buffer
	if err := customProfile.WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// CaptureSymbolProfile 采集Symbol Profile数据
func CaptureSymbolProfile() ([]byte, error) {
    var buf bytes.Buffer
    if err := pprof.Lookup("symbol").WriteTo(&buf, 0); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// CaptureContentionProfile 采集Contention Profile数据
func CaptureContentionProfile() ([]byte, error) {
    var buf bytes.Buffer
    if err := pprof.Lookup("contention").WriteTo(&buf, 0); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

//  CaptureSchedulerProfile 采集Scheduler Profile数据
func CaptureSchedulerProfile() ([]byte, error) {
    var buf bytes.Buffer
    if err := pprof.Lookup("sched").WriteTo(&buf, 0); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

//  CaptureGCProfile 采集GC Profile数据
func CaptureGCProfile() ([]byte, error) {
    var buf bytes.Buffer
    if err := pprof.Lookup("gc").WriteTo(&buf, 0); err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

