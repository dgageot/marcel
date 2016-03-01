package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dgageot/marcel/config"
	"github.com/dgageot/marcel/machine"
)

func executable(args []string) (string, []string) {
	switch {
	case len(args) == 0:
		return "docker", args
	case args[0] == "machine":
		return "docker-machine", args[1:]
	case args[0] == "compose":
		return "docker-compose", args[1:]
	default:
		return "docker", args
	}
}

func run(executable string, args []string) error {
	config, err := config.Load()
	if err != nil {
		return err
	}

	env, err := dockerEnv(config)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Env = append(os.Environ(), env...)

	return cmd.Run()
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
