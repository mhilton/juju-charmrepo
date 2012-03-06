package main

import (
	"fmt"
	"launchpad.net/juju/go/store"
	"launchpad.net/juju/go/log"
	stdlog "log"
	"net/http"
	"os"
)

func main() {
	log.Target = stdlog.New(os.Stdout, "", stdlog.LstdFlags)
	err := serve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func goodArg(arg string) bool {
	return len(arg) > 0 && arg[0] != '-'
}

func serve() error {
	if len(os.Args) != 3 || !goodArg(os.Args[1]) || !goodArg(os.Args[2]) {
		return fmt.Errorf("usage: csapi <mongo addr> <http addr>")
	}
	s, err := store.Open(os.Args[1])
	if err != nil {
		return err
	}
	defer s.Close()
	server, err := store.NewServer(s)
	if err != nil {
		return err
	}
	return http.ListenAndServe(os.Args[2], server)
}
