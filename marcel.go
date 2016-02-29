package main

import (
	"log"
	"os"

	"strings"

	"github.com/dgageot/marcel/config"
)

func marcel(args []string) error {
	switch {
	case len(args) == 2 && args[1] == "config":
		return config.Print()
	case len(args) == 3 && args[1] == "use" && args[2] == "local":
		return config.UseLocal()
	case len(args) == 3 && args[1] == "use" && !strings.HasPrefix(args[2], "tcp://"):
		return config.UseMachine(args[2])
	case len(args) == 3 && args[1] == "use":
		return config.UseUrl(args[2])
	case len(args) == 4 && args[1] == "use":
		return config.UseUrlWithTls(args[2], args[3])
	default:
		return run(executable(args...))
	}

	return nil
}

func main() {
	if err := marcel(os.Args); err != nil {
		log.Fatal(err)
	}
}
