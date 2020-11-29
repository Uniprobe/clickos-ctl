package xen

import "github.com/uniprobe/clickos-ctl/clickos"

//Domain represents a ClickOS Xen domain
type Domain struct {
	ID      uint
	Name    string
	XSPath  string
	Routers []clickos.Router
}
