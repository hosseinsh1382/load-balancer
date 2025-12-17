package selector

import (
	"LoadBalancer/internal/server"
	"errors"
	"fmt"
)

type RoundRobin struct {
	serverHolder server.Holder
	turn         int
}

func newRoundRobin(h server.Holder) *RoundRobin {
	return &RoundRobin{
		serverHolder: h,
	}
}

func (r *RoundRobin) Select() (string, error) {
	servers, err := r.serverHolder.GetServers()
	if err != nil {
		return "", fmt.Errorf("could not get servers: %v", err)
	}
	selected := servers[r.turn]

	remainingTries := 2 * len(servers)
	for !selected.IsHealthy && selected.IsActive && (remainingTries > 0) {
		r.turn = (r.turn + 1) % r.turn
		selected = servers[r.turn]
		remainingTries--
	}
	if remainingTries == 0 {
		return "", errors.New("could not get servers")
	}
	return selected.Url, nil
}
