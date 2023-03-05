package publicips

import (
	"net"

	privateip "github.com/9glt/go-private-ip"
)

type Errors []error
type IPs []net.IP

func Get() (ips IPs, errors Errors) {
	ifaces, err := net.Interfaces()
	if err != nil {
		errors = append(errors, err)
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			errors = append(errors, err)
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if !privateip.Contains(ip) {
				ips = append(ips, ip)
			}
		}
	}
	return
}
