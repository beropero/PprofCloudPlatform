package capture

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"time"

	"github.com/google/pprof/profile"
)

type Capture string

var ProfileSupportsDelta = map[Capture]bool{
	"allocs":       true,
	"block":        true,
	"goroutine":    true,
	"heap":         true,
	"mutex":        true,
	"threadcreate": true,
}

type CollectInput struct {
	gc     bool
	secStr string
	p      *pprof.Profile
}

func (name Capture) Collect(ctx context.Context, in CollectInput) (b []byte, err error) {
	var buf bytes.Buffer
	p := pprof.Lookup(string(name))
	in.p = p
	if p == nil {
		return nil, fmt.Errorf("unknown profile type: %s", name)
	}
	if in.secStr != "" {
		b, err = name.CollectDeltaProfile(ctx, in)
		return
	}
	if name == "heap" && in.gc {
		runtime.GC()
	}
	err = p.WriteTo(&buf, 0)
	return buf.Bytes(), err
}

func (name Capture) CollectDeltaProfile(ctx context.Context, in CollectInput) ([]byte, error) {
	var buf bytes.Buffer
	sec, err := strconv.ParseInt(in.secStr, 10, 64)
	if err != nil || sec <= 0 {
		return nil, fmt.Errorf("seconds parameter must be a positive integer")
	}
	if !ProfileSupportsDelta[name] {
		return nil, fmt.Errorf("delta profiles not supported for %s", name)
	}
	

	p0, err := collectProfile(in.p)
	if err != nil {
		return nil, fmt.Errorf("failed to collect profile")
	}

	// 添加间隔控制
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("profile canceled: %w", ctx.Err())
	case <-timer.C:
	}

	p1, err := collectProfile(in.p)
	if err != nil {
		return nil, fmt.Errorf("failed to collect profile")
	}
	ts := p1.TimeNanos
	dur := p1.TimeNanos - p0.TimeNanos

	p0.Scale(-1)

	p1, err = profile.Merge([]*profile.Profile{p0, p1})
	if err != nil {
		return nil, fmt.Errorf("failed to merge profiles")
	}

	p1.TimeNanos = ts
	p1.DurationNanos = dur
	err = p1.Write(&buf)
	return buf.Bytes(), err
}

func collectProfile(p *pprof.Profile) (*profile.Profile, error) {
	var buf bytes.Buffer
	if err := p.WriteTo(&buf, 0); err != nil {
		return nil, err
	}
	ts := time.Now().UnixNano()
	p0, err := profile.Parse(&buf)
	if err != nil {
		return nil, err
	}
	p0.TimeNanos = ts
	return p0, nil
}

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

//