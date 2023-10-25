// Package monitor provides a framework for building monitoring daemons.
package monitor

import (
	"context"
	"time"
)

// StaticDelay returns a static delay function for the specified period of time.
func StaticDelay(d time.Duration) func() {
	return func() { time.Sleep(d) }
}

// Monitor is a monitoring daemon that can periodically check some condition
// and, if it satisfies, trigger a callback and pause for a given delay.
type Monitor struct {
	ticker   *time.Ticker
	cond     func() bool
	callback func() error
	delay    func()
}

// NewMonitor creates a new monitoring daemon.
func NewMonitor(
	interval time.Duration,
	cond func() bool,
	callback func() error,
	delay func(),
) Monitor {
	return Monitor{
		ticker:   time.NewTicker(interval),
		cond:     cond,
		callback: callback,
		delay:    delay,
	}
}

// Run runs the monitoring daemon.
func (m Monitor) Run(ctx context.Context) error {
	defer m.ticker.Stop()

	for {
		select {
		case <-m.ticker.C:
			if m.cond() {
				if err := m.callback(); err != nil {
					return err
				}
				m.delay()
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
