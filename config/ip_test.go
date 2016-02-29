package config

import (
	"testing"

	"github.com/dgageot/marcel/machine"
	"github.com/stretchr/testify/assert"
)

type MockIPer struct {
	ipPerMachine map[string]string
}

func (m *MockIPer) IP(machine string) (string, error) {
	return m.ipPerMachine[machine], nil
}

func TestGetIP(t *testing.T) {
	defer func(iper func() machine.IPer) { machine.NewIPer = iper }(machine.NewIPer)
	machine.NewIPer = func() machine.IPer {
		return &MockIPer{
			ipPerMachine: map[string]string{
				"default": "192.168.99.100",
			},
		}
	}

	tests := []struct {
		config     *Config
		expectedIP string
	}{
		{&Config{
			Type: "local",
		}, "localhost"},
		{&Config{
			Type:    "machine",
			Machine: "default",
		}, "192.168.99.100"},
		{&Config{
			Type: "url",
			Url:  "tcp://192.168.99.150:2356",
		}, "192.168.99.150"},
		{&Config{
			Type: "url",
			Url:  "tcp://192.168.99.140",
		}, "192.168.99.140"},
	}

	for _, test := range tests {
		ip, err := getIP(test.config)

		assert.Equal(t, test.expectedIP, ip)
		assert.NoError(t, err)
	}
}
