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
	return &CliEnver{}
}

// CliEnver implements Enver by running docker-machine CLI.
type CliEnver struct{}

func (cli *CliEnver) Env(machine string) ([]string, error) {
	out, err := exec.Command("docker-machine", "env", "--shell=bash", machine).Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	return []string{removeExportAndQuotes(lines[0]), removeExportAndQuotes(lines[1]), removeExportAndQuotes(lines[2])}, nil
}

// removeExportAndQuotes converts
// export NAME="VALUE"
// to NAME=VALUE
func removeExportAndQuotes(line string) string {
	return strings.Replace(line[7:], `"`, "", -1)
}
