package wrap

type wrapinterface interface {
	CreateWrap(wrap *Wrap) error
	Getwrap(uuid string) (*Wrap, error)
}
