package common

type IDriver interface {
	//
	Resolve() (string, error)
}
