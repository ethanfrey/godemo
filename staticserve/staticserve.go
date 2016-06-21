package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	host := flag.String("host", "", "Hostname to bind to")
	port := flag.Int("port", 9000, "Port to serve on")
	flag.Parse()

	root := flag.Arg(0)
	if root == "" {
		fmt.Println(flag.Args())
		fmt.Println("Need one arg of a directory to serve")
		return
	}

	bind := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Serving %s on %s\n", root, bind)

	err := http.ListenAndServe(bind, http.FileServer(http.Dir(root)))
	if err != nil {
		panic(err)
	}
}
