package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	actions = map[string]string{
		"attache":     "attach",
		"construis":   "build",
		"engage":      "commit",
		"copie":       "cp",
		"crée":        "create",
		"changements": "diff",
		"évènements":  "events",
		"réalise":     "exec",
		"exporte":     "export",
		"historique":  "history",
		"images":      "images",
		"importe":     "import",
		"info":        "info",
		"inspecte":    "inspect",
		"tue":         "kill",
		"charge":      "load",
		"connecte":    "login",
		"déconnecte":  "logout",
		"journal":     "logs",
		"réseau":      "network",
		"bloque":      "pause",
		"port":        "port",
		"liste":       "ps",
		"récupère":    "pull",
		"publie":      "push",
		"renomme":     "rename",
		"redémarre":   "restart",
		"suppr":       "rm",
		"sim":         "rmi",
		"lance":       "run",
		"sauvegarde":  "save",
		"recherche":   "search",
		"démarre":     "start",
		"stats":       "stats",
		"arrête":      "stop",
		"labellise":   "tag",
		"premiers":    "top",
		"débloque":    "unpause",
		"version":     "version",
		"tome":        "volume",
		"attend":      "wait",
		"--aide":      "--help",
	}
)

func main() {
	args := os.Args

	newArgs := []string{}
	translateOutput := false

	if len(args) > 1 {
		action, present := actions[args[1]]
		if !present {
			log.Fatalf("marcel '%s' n'est pas une instruction reconnue.\nVoir 'marcel --aide'\n", args[1])
		}

		if len(args) > 2 && args[2] == "--aide" {
			newArgs = append([]string{action, "--help"}, args[3:]...)
			translateOutput = true
		} else {
			translateOutput = action == "--help"
			newArgs = append([]string{action}, args[2:]...)
		}
	}

	cmd := exec.Command("docker", newArgs...)
	if translateOutput {
		output, err := cmd.CombinedOutput()
		if err != nil {
			os.Exit(1)
		}

		help := string(output)
		help = strings.Replace(help, "Usage:\tdocker ", "Utilisation\tmarcel ", -1)
		for fr, us := range actions {
			help = strings.Replace(help, us, fr, -1)
		}

		fmt.Print(help)
	} else {
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			os.Exit(1)
		}
	}
}
