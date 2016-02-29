package main

import (
	"log"
	"os"
	"os/exec"

	"strings"

	"fmt"

	"github.com/dgageot/marcel/config"
	"github.com/dgageot/marcel/machine"
)

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

func dockerEnv(config *config.Config) ([]string, error) {
	switch config.Type {
	case "local":
		return []string{"DOCKER_TLS_VERIFY=", "DOCKER_HOST=", "DOCKER_CERT_PATH="}, nil
	case "machine":
		env, err := machine.NewEnver().Env(config.Machine)
		if err != nil {
			return nil, err
		}

		return env, nil
	case "url":
		if config.CertPath != "" {
			return []string{"DOCKER_TLS_VERIFY=1", "DOCKER_HOST=" + config.Url, "DOCKER_CERT_PATH=" + config.CertPath}, nil
		} else {
			return []string{"DOCKER_TLS_VERIFY=", "DOCKER_HOST=" + config.Url, "DOCKER_CERT_PATH="}, nil
		}
	}

	return nil, fmt.Errorf("Unknown type: %s", config.Type)
}

func runCommand(executable string, args []string) {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	env, err := dockerEnv(config)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(executable, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = append(os.Environ(), env...)

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
			Type: "url",
			Url:  args[2],
		})
	case len(args) == 4 && args[1] == "use":
		config.Save(&config.Config{
			Type:     "url",
			Url:      args[2],
			CertPath: args[3],
		})
	default:
		runCommand(findCommand(args...))
	}
}
