package models

type ReturnModel[T any] struct {
	status  int
	data    T
	message string
}
