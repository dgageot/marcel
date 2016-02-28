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

func TestMachineVersion(t *testing.T) {
	executable, args := findCommand("marcel", "machine", "--version")

	assert.Equal(t, "docker-machine", executable)
	assert.Equal(t, []string{"--version"}, args)
}
