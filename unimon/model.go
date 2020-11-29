package unimon

//UnimonElement represents a Unimon Element in a ClickOS router
type UnimonElement struct {
	Name   string
	XSPath string
	Data   []UnimonData
}

//UnimonData represents a single data point from a Unimon element
type UnimonData struct {
	timestamp uint64
	data      interface{}
}
