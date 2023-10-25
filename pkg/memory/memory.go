// Package memory provides memory monitoring tools.
package memory

import (
	"fmt"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/lzakharov/autopprof/pkg/monitor"
	"github.com/lzakharov/autopprof/pkg/storage"
)

const (
	defaultInterval   = 5 * time.Second
	defaultDelay      = time.Minute
	defaultTimeFormat = "20060102T150405"
)

func defaultName() string {
	return time.Now().Format(defaultTimeFormat) + "_profile.pb.gz"
}

// Option is a heap monitoring daemon option.
type Option func(*params)

// WithInterval sets the monitoring interval.
func WithInterval(interval time.Duration) Option {
	return func(p *params) {
		p.interval = interval
	}
}

// WithStorage sets the storage for heap profiles.
func WithStorage(storage storage.Storage) Option {
	return func(p *params) {
		p.storage = storage
	}
}

// WithNameFunc sets the heap profile file name function.
func WithNameFunc(f func() string) Option {
	return func(p *params) {
		p.name = f
	}
}

// WithDelay sets the delay function.
func WithDelay(delay func()) Option {
	return func(p *params) {
		p.delay = delay
	}
}

type params struct {
	interval time.Duration
	storage  storage.Storage
	name     func() string
	delay    func()
}

// NewMonitor creates a new memory monitoring daemon. This daemon periodically
// checks the amount of memory used and, if it exceeds a specified threshold (in
// bytes), dumps heap profile. By default, the daemon checks every 5 seconds,
// writes profiles to the current folder, and falls asleep for a minute after
// triggering. See defaultInterval, defaultDelay and defaultName. You can change
// the default settings using options.
func NewMonitor(threshold uint64, opts ...Option) monitor.Monitor {
	params := params{
		interval: defaultInterval,
		storage:  storage.FS{},
		name:     defaultName,
		delay:    monitor.StaticDelay(defaultDelay),
	}

	for _, opt := range opts {
		opt(&params)
	}

	return monitor.NewMonitor(
		params.interval,
		memoryThreshold(threshold),
		DumpHeapProfile(params.storage, params.name),
		params.delay)
}

func memoryThreshold(threshold uint64) func() bool {
	return func() bool {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)
		return stats.Sys > threshold
	}
}

// DumpHeapProfile creates a function that dumps heap profile into the given
// storage, generating a filename from the given name function.
func DumpHeapProfile(storage storage.Storage, name func() string) func() error {
	return func() error {
		f, err := storage.Open(name())
		if err != nil {
			return fmt.Errorf("open file: %w", err)
		}
		defer f.Close()

		if err := pprof.WriteHeapProfile(f); err != nil {
			return fmt.Errorf("write heap profile: %w", err)
		}

		return nil
	}
}
