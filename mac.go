package mac

import (
	"errors"
	"net"
	"regexp"
)

const (
	macRegex  = `(?i)^([a-z0-9]{1,2}[:-]){5}[a-z0-9]{1,2}$`
	// 00:00:00:00:00:00
	// 00-00-00-00-00-00
	zeroRegex = `([0]{1,2}[:-]){5}[0]{1,2}`
)

func getMac(iface string) (string, error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return "", err
	}

	regex := regexp.MustCompile(zeroRegex)

	if iface != "" {
		for _, v := range interfaces {
			addr := v.HardwareAddr.String()

			if v.Name == iface {
				if regex.MatchString(addr) == false {
					return addr, nil
				}
				break
			}
		}
	} else {
		for _, v := range interfaces {
			addr := v.HardwareAddr.String()

			if addr == "" {
				continue
			}

			if regex.MatchString(addr) == false {
				return addr, nil
			}
		}
	}

	return "", errors.New("mac: Not Found")
}

func AddrWithIface(iface string) (string, error) {
	return getMac(iface)
}

func Default() (string, error) {
	return getMac("")
}

func IsMAC(addr string) bool {
	regex := regexp.MustCompile(macRegex)
	return regex.MatchString(addr)
}
