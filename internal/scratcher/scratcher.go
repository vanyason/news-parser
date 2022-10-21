package scratcher

type Scratcher interface {
	Scratch() (res Result, err error)
}
