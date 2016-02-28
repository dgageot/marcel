package main

import (
	"os"
	"os/exec"
)

// marcel
// marcel [run|build|...] [...]
// marcel machine [...]
// marcel compose [...]
// marcel use machine default
// marcel use docker local
// marcel use docker url certs
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
	runCommand(findCommand(os.Args...))
}
