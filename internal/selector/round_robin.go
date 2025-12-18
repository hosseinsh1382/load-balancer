package selector

import (
	"LoadBalancer/internal/server"
	"errors"
	"fmt"
	"log"
)

type RoundRobin struct {
	serverHolder server.Holder
	turn         int
}

func NewRoundRobin(h server.Holder) *RoundRobin {
	return &RoundRobin{
		serverHolder: h,
	}
}

func (r *RoundRobin) Select() (string, error) {
	log.Println("[ProxyHandler] Select")
	servers, err := r.serverHolder.Servers()
	if err != nil {
		return "", fmt.Errorf("could not get servers: %v", err)
	}
	selected := servers[r.turn]

	remainingTries := 2 * len(servers)
	for !(selected.IsHealthy && selected.IsActive) && (remainingTries > 0) {
		r.turn = (r.turn + 1) % len(servers)
		selected = servers[r.turn]
		remainingTries--
	}
	if remainingTries == 0 {
		return "", errors.New("could not get servers")
	}
	r.turn = (r.turn + 1) % len(servers)
	return selected.Url, nil
}
