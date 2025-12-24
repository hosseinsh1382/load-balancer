package server

import (
	"LoadBalancer/pkg"
	"sync"
)

type JsonHolder struct {
	jsonReader    pkg.JsonReader[[]Registry]
	healthChecker HealthChecker
	servers       []*Server
	locker        sync.RWMutex
}

func NewJsonHolder(reader pkg.JsonReader[[]Registry], h HealthChecker) *JsonHolder {
	j := &JsonHolder{
		jsonReader:    reader,
		healthChecker: h,
	}
	err := j.UpdateServers()
	if err != nil {
		panic(err)
	}
	return j
}

func (j *JsonHolder) UpdateServers() error {
	registries, err := j.jsonReader.ReadJson()
	j.locker.Lock()
	if err != nil {
		return err
	}
	for _, r := range registries {
		health := j.healthChecker.CheckHealth(r.Url)
		s := Server{
			Name:      r.Name,
			IsActive:  r.IsActive,
			IsHealthy: health,
			Url:       r.Url,
		}
		j.servers = append(j.servers, &s)
	}
	j.locker.Unlock()
	return nil
}
func (j *JsonHolder) Servers() ([]*Server, error) {
	return j.servers, nil
}
