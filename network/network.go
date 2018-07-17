package network

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

type NTconfig struct {
	Bridge string
	Vehost string
	Vepeer string
	Mtu    int
	Ipaddr string
}

func (ne *NTconfig) attach() error {
	brl, err := netlink.LinkByName(ne.Bridge)
	if err != nil {
		return err
	}
	br, ok := brl.(*netlink.Bridge)
	if !ok {
		return fmt.Errorf("Wrong device type %T", brl)
	}
	host, err := netlink.LinkByName(ne.Vehost)
	if err != nil {
		return fmt.Errorf("Cannot link %s: %v ", ne.Vehost, err)
	}

	if err := netlink.LinkSetMaster(host, br); err != nil {
		return fmt.Errorf("Failed to link Master %s %s: %v ", host, br, err)
	}
	if err := netlink.LinkSetMTU(host, ne.Mtu); err != nil {
		return fmt.Errorf("Failed to set MTU %d: %v ", ne.Mtu, err)
	}
	if err := netlink.LinkSetUp(host); err != nil {
		return fmt.Errorf("Failed to link up %s: %v ", host, err)
	}
	return nil
}

func (ne *NTconfig) Create(nspid int)(err error){
	veth := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{
			Name: ne.Vehost,
		},
		PeerName: ne.Vepeer,
	}
	if err := netlink.LinkAdd(veth); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			netlink.LinkDel(veth)
		}
	}()
	if err := ne.attach(); err != nil {
		return err
	}
	child, err := netlink.LinkByName(ne.Vepeer)
	if err != nil {
		return err
	}
	if err = netlink.LinkSetNsPid(child, nspid); err != nil {
		return err
	}
	return nil
}

func detach(vehostname string) (err error) {
	return netlink.LinkSetMaster(&netlink.Device{LinkAttrs: netlink.LinkAttrs{Name: vehostname}}, nil)
}

func (ne *NTconfig) Remove() (err error) {
	netlink.LinkDel(&netlink.Device{LinkAttrs: netlink.LinkAttrs{Name: ne.Vehost}})
	return nil
}
