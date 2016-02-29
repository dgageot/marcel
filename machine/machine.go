package machine

import (
	"os/exec"
	"strings"
)

type Enver interface {
	Env(machine string) ([]string, error)
}

type CliEnver struct{}

var NewEnver = func() Enver {
	return &CliEnver{}
}

func (cli *CliEnver) Env(machine string) ([]string, error) {
	out, err := exec.Command("docker-machine", "env", "--shell=bash", machine).CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	return []string{strings.Replace(lines[0][7:], `"`, "", -1), strings.Replace(lines[1][7:], `"`, "", -1), strings.Replace(lines[2][7:], `"`, "", -1)}, nil
}
