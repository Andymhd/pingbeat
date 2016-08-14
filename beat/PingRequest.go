package pingbeat

import (
	"github.com/elastic/beats/libbeat/logp"
	"golang.org/x/net/icmp"
	"net"
	"os"
)

type PingRequest struct {
	text_payload   *icmp.Message
	binary_payload []byte
	ping_type      icmp.Type
	target         *net.IPAddr
	iface          string
}

func (request *PingRequest) Encode(seq_no int, ping_type icmp.Type, target string, iface string) {
	request.iface = iface
	request.target = &net.IPAddr{IP: net.ParseIP(target), Zone: iface}
	request.ping_type = ping_type
	request.text_payload = &icmp.Message{
		Type: request.ping_type, Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  seq_no,
			Data: []byte("pingbeat: y'know, for pings"),
		},
	}
	binary, err := request.text_payload.Marshal(nil)
	if err != nil {
		logp.Err("Error encoding packet: %v", err)
	} else {
		request.binary_payload = binary
	}
}
