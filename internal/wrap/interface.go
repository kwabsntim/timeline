package wrap

type wrapinterface interface {
	CreateWrap(wrap *Wrap) error
	Getwrap(uuid string) (*Wrap, error)
}
type ServiceInterface interface {
	MakeWrap(name string) (*Wrap, error)
	GetWrap(uuid string) (*Wrap, error)
}
