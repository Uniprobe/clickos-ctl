package xenstore

import (
	"fmt"

	xs "github.com/joelnb/xenstore-go"
)

const (
	XSSocketPathDefault string = "/var/run/xenstored/socket"
	basePath            string = "/local/domain"
)

var (
	xenstoreClient *xs.Client
)

func CreateClient(path string) error {
	xsc, err := xs.NewUnixSocketClient(path)
	xenstoreClient = xsc
	return err
}

func GetPathChildren(path string) ([]string, error) {
	if xenstoreClient == nil {
		return nil, fmt.Errorf("Xenstore client not yet created")
	}
	return xenstoreClient.List(basePath + path)
}

func GetPathContent(path string) (string, error) {
	if xenstoreClient == nil {
		return "", fmt.Errorf("Xenstore client not yet created")
	}
	return xenstoreClient.Read(basePath + path)
}

func GetDomainPath(domid int) (string, error) {
	if xenstoreClient == nil {
		return "", fmt.Errorf("Xenstore client not yet created")
	}
	return xenstoreClient.GetDomainPath(domid)
}
