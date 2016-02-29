package main

import (
	"os"
	"testing"

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
