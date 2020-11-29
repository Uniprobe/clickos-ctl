package clickos

const (
	runningStateString string = "Running"
	haltStateString    string = "Halted"
	errorStateString   string = "Error"
)

type Domain struct {
	ID      uint
	Name    string
	XSPath  string
	Routers []Router
}

type Router struct {
	ID             uint
	XSPath         string
	UnimonElements []UnimonElement
}

type UnimonElement struct {
	Name string
	Data []UnimonData
}

type UnimonData struct {
	timestamp uint64
	data      interface{}
}
