package main

import (
	"os"
	"os/exec"

	"strings"

	"github.com/dgageot/marcel/config"
)

type Config struct {
	Type      string
	Machine   string
	Url       string
	TlsVerify bool
	CertPath  string
}

// [X] marcel
// [X] marcel [run|build|...] [...]
// [X] marcel machine [...]
// [X] marcel compose [...]
// [X] marcel use machine default
// [X] marcel use local
// [X] marcel use tcp://192.168.99.100:2376 [~/.docker/certs]"
// [ ] Pass config to docker
// [ ] Pass config to docker-compose
//
func findCommand(args ...string) (string, []string) {
	switch {
	case len(args) < 2:
		return "docker", args[1:]
	case args[1] == "machine":
		return "docker-machine", args[2:]
	case args[1] == "compose":
		return "docker-compose", args[2:]
	default:
		return "docker", args[1:]
	}
}

func runCommand(executable string, args []string) {
	cmd := exec.Command(executable, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

func main() {
	args := os.Args

	switch {
	case len(args) == 3 && args[1] == "use" && args[2] == "local":
		config.Save(&config.Config{
			Type: "local",
		})
	case len(args) == 3 && args[1] == "use" && !strings.HasPrefix(args[2], "tcp://"):
		config.Save(&config.Config{
			Type:    "machine",
			Machine: args[2],
		})
	case len(args) == 3 && args[1] == "use":
		config.Save(&config.Config{
			Type:      "url",
			Url:       args[2],
			TlsVerify: false,
		})
	case len(args) == 4 && args[1] == "use":
		config.Save(&config.Config{
			Type:      "url",
			Url:       args[2],
			TlsVerify: true,
			CertPath:  args[3],
		})
	default:
		runCommand(findCommand(args...))
	}
}
