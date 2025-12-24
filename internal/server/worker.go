package server

import (
	"fmt"
	"time"
)

type HealthCheckerWorker struct {
	serverHolder  Holder
	healthChecker HealthChecker
	duration      time.Duration
}

func NewHealthWorker(serverHolder Holder, healthChecker HealthChecker, d time.Duration) *HealthCheckerWorker {
	return &HealthCheckerWorker{
		serverHolder:  serverHolder,
		healthChecker: healthChecker,
		duration:      d,
	}
}

func (a *HealthCheckerWorker) Start() {
	ticker := time.NewTicker(a.duration)
	start := time.Now()
	for {
		select {
		case <-ticker.C:
			servers, err := a.serverHolder.Servers()
			if err != nil {
				panic(err)
			}
			for _, server := range servers {
				status := a.healthChecker.CheckHealth(server.Url)
				server.IsHealthy = status
			}
		default:
			elapsed := time.Since(start).Round(time.Millisecond)
			fmt.Println("[HealthCheckerWorker] elapsed:", elapsed)
			time.Sleep(1 * time.Second)
		}
	}
}
