package server

import "LoadBalancer/pkg"

type JsonHolder struct {
	jsonReader pkg.JsonReader[[]Registry]
	servers    []Server
}

func NewJsonHolder(reader pkg.JsonReader[[]Registry]) *JsonHolder {
	j := &JsonHolder{
		jsonReader: reader,
	}
	err := j.UpdateServers()
	if err != nil {
		panic(err)
	}
	return j
}

func (j *JsonHolder) UpdateServers() error {
	registries, err := j.jsonReader.ReadJson()
	if err != nil {
		return err
	}
	for _, r := range registries {
		s := Server{
			Name:      r.Name,
			IsActive:  r.IsActive,
			IsHealthy: true,
			Url:       r.Url,
		}
		j.servers = append(j.servers, s)
	}
	return nil
}
func (j *JsonHolder) Servers() ([]Server, error) {
	return j.servers, nil
}
