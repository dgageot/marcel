package main

import (
	"log"
	"os"

	"strings"

	"github.com/dgageot/marcel/config"
)

func marcel(args []string) error {
	switch {
	case len(args) == 1 && args[0] == "ip":
		return config.PrintIP()
	case len(args) == 2 && args[0] == "use" && args[1] == "local":
		return config.UseLocal()
	case len(args) == 2 && args[0] == "use" && !strings.HasPrefix(args[1], "tcp://"):
		return config.UseMachine(args[1])
	case len(args) == 2 && args[0] == "use":
		return config.UseUrl(args[1])
	case len(args) == 3 && args[0] == "use":
		return config.UseUrlWithTls(args[1], args[2])
	default:
		return run(executable(args))
	}

	return nil
}

func main() {
	if err := marcel(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
