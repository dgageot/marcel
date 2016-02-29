package main

import (
	"os"
	"testing"

	"github.com/dgageot/marcel/config"
	"github.com/stretchr/testify/assert"
)

func ExampleDockerVersion() {
	os.Args = []string{"marcel", "--version"}

	main()

	// Output:
	// Docker version 1.10.1, build 9e83765
}

func TestEmptyArgs(t *testing.T) {
	executable, args := findCommand("marcel")

	assert.Equal(t, "docker", executable)
	assert.Equal(t, []string{}, args)
}

func TestFindCommand(t *testing.T) {
	tests := []struct {
		args               []string
		expectedExecutable string
		expectedArgs       []string
	}{
		{[]string{"marcel"}, "docker", []string{}},
		{[]string{"marcel", "--version"}, "docker", []string{"--version"}},
		{[]string{"marcel", "machine", "--version"}, "docker-machine", []string{"--version"}},
		{[]string{"marcel", "compose", "--version"}, "docker-compose", []string{"--version"}},
		{[]string{"marcel", "run", "hello-world"}, "docker", []string{"run", "hello-world"}},
	}

	for _, test := range tests {
		executable, args := findCommand(test.args...)

		assert.Equal(t, test.expectedExecutable, executable)
		assert.Equal(t, test.expectedArgs, args)

	}
}

func TestUseMachine(t *testing.T) {
	tests := []struct {
		args           []string
		expectedConfig config.Config
	}{
		{[]string{"default"}, config.Config{
			Type:    "machine",
			Machine: "default",
		}},
		{[]string{"other"}, config.Config{
			Type:    "machine",
			Machine: "other",
		}},
		{[]string{"local"}, config.Config{
			Type: "local",
		}},
		{[]string{"tcp://192.168.99.100:2376", "~/.docker/certs"}, config.Config{
			Type:      "url",
			Url:       "tcp://192.168.99.100:2376",
			TlsVerify: true,
			CertPath:  "~/.docker/certs",
		}},
		{[]string{"tcp://192.168.99.150:2376"}, config.Config{
			Type:      "url",
			Url:       "tcp://192.168.99.150:2376",
			TlsVerify: false,
		}},
	}

	for _, test := range tests {
		os.Args = append([]string{"marcel", "use"}, test.args...)

		main()

		config, err := config.Load()

		assert.Equal(t, test.expectedConfig, *config)
		assert.NoError(t, err)
	}
}
