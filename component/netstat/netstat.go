package netstat

import (
	"fmt"
	"net"
	"strconv"
)

// SockAddr represents an ip:port pair
type SockAddr struct {
	IP   net.IP
	Port uint16
}

func (s *SockAddr) String() string {
	return fmt.Sprintf("%v:%d", s.IP, s.Port)
}

// SockTabEntry type represents each line of the /proc/net/[tcp|udp]
type SockTabEntry struct {
	ino        string
	LocalAddr  *SockAddr
	RemoteAddr *SockAddr
	State      SkState
	UID        uint32
	Process    *Process
}

// Process holds the PID and process name to which each socket belongs
type Process struct {
	Pid  int
	Name string
}

func (p *Process) String() string {
	if p != nil {
		return fmt.Sprintf("%d/%s", p.Pid, p.Name)
	}
	return "-/-"
}

type AcceptFn func(*SockTabEntry) bool

// SkState type represents socket connection state
type SkState uint8

func (s SkState) String() string {
	return skStates[s]
}

// LookupPort searches among system connections by ip address plus port. gives process id/process name as return.
func LookupPort(host string, port string) string {
	if host == "" || port == "" {
		return "-/-"
	}
	p, err := strconv.ParseUint(port, 0, 0)
	if err != nil {
		return "-/-"
	}
	connections, err := osSocks(func(s *SockTabEntry) bool {
		if s.LocalAddr.IP.String() == host && uint16(p) == s.LocalAddr.Port {
			return true
		}
		return false
	})
	if err != nil || len(connections) < 1 {
		return "empty"
	}
	return connections[0].Process.String()
}
