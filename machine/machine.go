package machine

import (
	"os/exec"
	"strings"
)

// Enver provides the env variables for a machine name.
type Enver interface {
	Env(machine string) ([]string, error)
}

// NewEnver creates a default Enver implementation.
var NewEnver = func() Enver {
	return &MachineCLI{}
}

// IPer provides the IP address for a machine name.
type IPer interface {
	IP(machine string) (string, error)
}

// NewIPer creates a default IPer implementation.
var NewIPer = func() IPer {
	return &MachineCLI{}
}

// MachineCLI implements Enver and IPer by running docker-machine CLI.
type MachineCLI struct{}

func (cli *MachineCLI) Env(machine string) ([]string, error) {
	out, err := exec.Command("docker-machine", "env", "--shell=bash", machine).Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	return []string{removeExportAndQuotes(lines[0]), removeExportAndQuotes(lines[1]), removeExportAndQuotes(lines[2])}, nil
}

func (cli *MachineCLI) IP(machine string) (string, error) {
	out, err := exec.Command("docker-machine", "ip", machine).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// removeExportAndQuotes converts
// export NAME="VALUE"
// to NAME=VALUE
func removeExportAndQuotes(line string) string {
	return strings.Replace(line[7:], `"`, "", -1)
}
