package selector

type Selector interface {
	Select() (string, error)
}
