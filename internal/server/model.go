package server

type Server struct {
	Name      string
	IsActive  bool
	IsHealthy bool
	Url       string
}

type Registry struct {
	Name     string
	Url      string
	IsActive bool
}
