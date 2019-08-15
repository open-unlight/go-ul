package protocol

// Callback ...
type Callback interface {
	Arguments() []string
	Invoke(...interface{}) interface{}
}
