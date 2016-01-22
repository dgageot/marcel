package main

import "os"

func ExampleHello() {
	os.Args = []string{"marcel", "attache", "--aide"}

	main()

	// Output:
	// Utilisation:	marcel attache [OPTIONS] CONTAINER
	//
	// Attach to a running container
	//
	//   --help=false        Print usage
	//   --no-stdin=false    Do not attache STDIN
	//   --sig-proxy=true    Proxy all received signals to the process
}
