package dhcp

import (
	"net"

	"github.com/insomniacslk/dhcp/dhcpv4"
)

func (r *Role) handleDHCPRelease4(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	// TODO: Implement DHCP release
}
