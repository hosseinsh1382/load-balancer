package server

import "LoadBalancer/pkg"

type JsonHolder struct {
	jsonReader pkg.JsonReader[[]Server]
	servers    []Server
}

func NewJsonHolder(reader pkg.JsonReader[[]Server]) *JsonHolder {
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
	s, err := j.jsonReader.ReadJson()
	if err != nil {
		return err
	}
	j.servers = s
	return nil
}
func (j *JsonHolder) Servers() ([]Server, error) {
	return j.servers, nil
}
