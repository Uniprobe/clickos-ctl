package clickos

import "github.com/uniprobe/clickos-ctl/unimon"

const (
	runningStateString string = "Running"
	haltStateString    string = "Halted"
	errorStateString   string = "Error"
)

//Router represents a single ClickOS router
type Router struct {
	ID             uint
	XSPath         string
	UnimonElements []unimon.UnimonElement
}
