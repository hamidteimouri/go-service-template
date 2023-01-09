package main

import (
	"goservicetemplate/cmd/bootstrap"
	_ "goservicetemplate/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Start the application
	bootstrap.Start()

	// wait for `Ctrl+c` or docker stop/restart signal
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
	<-ch

	// Stop the application
	bootstrap.Stop()
}
