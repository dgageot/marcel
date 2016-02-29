package main

import (
	"os"
	"testing"

	"github.com/dgageot/marcel/config"
	"github.com/dgageot/marcel/machine"
	"github.com/stretchr/testify/assert"
)

func ExampleDockerVersion() {
	os.Args = []string{"marcel", "--version"}

	main()

	// Output:
	// Docker version 1.10.1, build 9e83765
}

func TestEmptyArgs(t *testing.T) {
	executable, args := executable("marcel")

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
		executable, args := executable(test.args...)

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
			Type:     "url",
			Url:      "tcp://192.168.99.100:2376",
			CertPath: "~/.docker/certs",
		}},
		{[]string{"tcp://192.168.99.150:2376"}, config.Config{
			Type: "url",
			Url:  "tcp://192.168.99.150:2376",
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

type MockEnver struct {
	envPerMachine map[string][]string
}

func (m *MockEnver) Env(machine string) ([]string, error) {
	return m.envPerMachine[machine], nil
}

func TestDockerEnv(t *testing.T) {
	defer func(enver func() machine.Enver) { machine.NewEnver = enver }(machine.NewEnver)
	machine.NewEnver = func() machine.Enver {
		return &MockEnver{
			envPerMachine: map[string][]string{
				"default": []string{"DOCKER_TLS_VERIFY=1", "DOCKER_HOST=tcp://192.168.99.100:2376", "DOCKER_CERT_PATH=/Users/dgageot/.docker/machine/machines/default"},
			},
		}
	}

	tests := []struct {
		config      *config.Config
		expectedEnv []string
	}{
		{&config.Config{
			Type:    "machine",
			Machine: "default",
		}, []string{"DOCKER_TLS_VERIFY=1", "DOCKER_HOST=tcp://192.168.99.100:2376", "DOCKER_CERT_PATH=/Users/dgageot/.docker/machine/machines/default"}},
		{&config.Config{
			Type: "local",
		}, []string{"DOCKER_TLS_VERIFY=", "DOCKER_HOST=", "DOCKER_CERT_PATH="}},
		{&config.Config{
			Type: "url",
			Url:  "tcp://192.168.99.150:2376",
		}, []string{"DOCKER_TLS_VERIFY=", "DOCKER_HOST=tcp://192.168.99.150:2376", "DOCKER_CERT_PATH="}},
		{&config.Config{
			Type:     "url",
			Url:      "tcp://192.168.99.160:2376",
			CertPath: "/certs",
		}, []string{"DOCKER_TLS_VERIFY=1", "DOCKER_HOST=tcp://192.168.99.160:2376", "DOCKER_CERT_PATH=/certs"}},
	}

	for _, test := range tests {
		env, err := dockerEnv(test.config)

		assert.Equal(t, test.expectedEnv, env)
		assert.NoError(t, err)
	}
}
