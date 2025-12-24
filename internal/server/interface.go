package server

type Holder interface {
	Servers() ([]*Server, error)
	UpdateServers() error
}

type HealthChecker interface {
	CheckHealth(url string) bool
}
