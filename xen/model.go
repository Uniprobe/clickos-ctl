package xen

import (
	"github.com/uniprobe/clickos-ctl/clickos"
	"github.com/uniprobe/clickos-ctl/xenstore"
)

//Domain represents a ClickOS Xen domain
type Domain struct {
	ID      int
	Name    string
	Routers []clickos.Router
}

func (d *Domain) XSPath() (string, error) {
	return xenstore.GetDomainPath(d.ID)
}
