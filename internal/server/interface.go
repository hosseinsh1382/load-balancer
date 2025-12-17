package server

type Holder interface {
	GetServers() ([]Server, error)
	UpdateServer() error
}
