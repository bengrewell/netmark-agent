package core

type Module interface {
	Run() (value map[string]interface{})
}
