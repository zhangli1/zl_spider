package model

type Model interface {
	Run() interface{}
	Destruct(interface{}) interface{}
}
