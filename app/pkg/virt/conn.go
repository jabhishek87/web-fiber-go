package virt

import (
	"log"

	"libvirt.org/go/libvirt"
)

// Connection ...
type Connection struct {
	Host       string
	Connection *libvirt.Connect
}

// VMInfo ...
type VMInfo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// New ...
func New(host string) Connection {
	return Connection{host, nil}
}

func (vc *Connection) connect() {
	conn, err := libvirt.NewConnect(vc.Host)
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}
	vc.Connection = conn
}

// ListAllVM ...
func (vc *Connection) ListAllVM() ([]VMInfo, error) {
	vc.connect()
	defer vc.Connection.Close()
	domains, err := vc.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_PERSISTENT)
	if err != nil {
		return nil, err
	}
	vmInfoList := []VMInfo{}
	for _, vm := range domains {
		vmID, _ := vm.GetID()
		vmName, _ := vm.GetName()
		vmState, _ := vm.IsActive()
		vmInfo := VMInfo{vmID, vmName, b2str(vmState)}
		vmInfoList = append(vmInfoList, vmInfo)
	}
	return vmInfoList, err
}

func b2str(b bool) string {
	if b {
		return "On"
	}
	return "Off"
}
