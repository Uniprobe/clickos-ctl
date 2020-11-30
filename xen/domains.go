package xen

import (
	"fmt"
	"strconv"

	"github.com/uniprobe/clickos-ctl/xenstore"
)

// GetClickOSDomains returns the dom ids and names of any clickos domain found
func GetClickOSDomains() ([]Domain, error) {
	children, err := xenstore.GetPathChildren("")
	if err != nil {
		fmt.Printf("🆘 Failed to get domain list from xenstore")
		return nil, err
	}
	domains := make([]Domain, 0, len(children))
	for _, domid := range children {
		name, err := xenstore.GetPathContent("/" + domid + "/name")
		if err != nil {
			fmt.Println("⚠️ Failed to get domain name")
			name = "unknown name"
		}
		id, err := strconv.Atoi(domid)
		if err != nil {
			fmt.Println("⚠️ Domain ID could not be parsed correctly")
			id = -1
		}
		domains = append(domains, Domain{
			ID:   id,
			Name: name,
		})
	}
	return domains, nil
}

func GetDomainByName(name string) {

}

func GetDomainByID(ID uint) {

}

func getDomain(ID uint) {

}

func (domain *Domain) InstallRouter(configPath string) (uint, error) {
	return 0, nil
}

func (domain *Domain) RemoveRouter(rid uint) error {
	return nil
}
