package henceforth

type Completable interface {
	Complete() error
	IsComplete() bool
}
