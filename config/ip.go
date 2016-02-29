package config

import (
	"fmt"

	"strings"

	"github.com/dgageot/marcel/machine"
)

func getIP(current *Config) (string, error) {
	switch current.Type {
	case "local":
		return "localhost", nil
	case "machine":
		ip, err := machine.NewIPer().IP(current.Machine)
		if err != nil {
			return "", err
		}

		return ip, nil
	case "url":
		return ipForUrl(current.Url), nil
	}

	return "", fmt.Errorf("Unknown type: %s", current.Type)
}

func ipForUrl(url string) string {
	ipWithPort := strings.SplitAfter(url, "://")[1]
	ip := strings.Split(ipWithPort, ":")[0]

	return ip
}
