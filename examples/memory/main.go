package main

import (
	"context"
	"strconv"
	"time"

	"github.com/lzakharov/autopprof/pkg/memory"
)

const threshold = 1 << 20 // 1Mb

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	monitor := memory.NewMonitor(threshold)
	go monitor.Run(ctx)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = allocate()
		}
	}
}

func allocate() map[int]string {
	m := make(map[int]string)
	for i := 0; i < 1024; i++ {
		m[i] = strconv.Itoa(i)
	}
	return m
}
