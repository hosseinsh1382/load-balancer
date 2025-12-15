package selector

type StaticSelector struct {
}

func NewStaticSelector() *StaticSelector {
	return &StaticSelector{}
}

func (s *StaticSelector) Select() (string, error) {
	return "http://localhost:8084", nil
}
