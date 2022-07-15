package dhcp

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
	log "github.com/sirupsen/logrus"
)

func (r *Role) recoverMiddleware4(inner server4.Handler) server4.Handler {
	return func(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			if e, ok := err.(error); ok {
				r.log.WithError(e).Warning("recover in dhcp handler")
				sentry.CaptureException(e)
			} else {
				r.log.WithField("panic", err).Warning("recover in dhcp handler")
			}
		}()
		inner(conn, peer, m)
	}
}

func (r *Role) logDHCPMessage(m *dhcpv4.DHCPv4, fields log.Fields) {
	f := log.Fields{
		"deviceIdentifier": r.DeviceIdentifier(m),
		"opCode":           m.OpCode.String(),
		"hopCount":         m.HopCount,
		"transactionID":    m.TransactionID.String(),
		"flagsToString":    m.FlagsToString(),
		"clientIPAddr":     m.ClientIPAddr.String(),
		"yourIPAddr":       m.YourIPAddr.String(),
		"serverIPAddr":     m.ServerIPAddr.String(),
		"gatewayIPAddr":    m.GatewayIPAddr.String(),
		"hostname":         m.HostName(),
		"clientIdentifier": hex.EncodeToString(m.Options.Get(dhcpv4.OptionClientIdentifier)),
	}
	r.log.WithFields(f).WithFields(fields).Info(m.MessageType().String())
}

func (r *Role) loggingMiddleware4(inner server4.Handler) server4.Handler {
	return func(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
		start := time.Now()
		inner(conn, peer, m)
		duration := float64(time.Since(start)) / float64(time.Millisecond)
		f := log.Fields{
			"client":    peer.String(),
			"localAddr": conn.LocalAddr().String(),
			"runtimeMS": fmt.Sprintf("%0.3f", duration),
		}
		r.logDHCPMessage(m, f)
	}
}
