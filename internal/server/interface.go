package server

type Holder interface {
	Servers() ([]Server, error)
	UpdateServers() error
}
