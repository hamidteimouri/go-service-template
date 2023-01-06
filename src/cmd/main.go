package main

import (
	"goservicetemplate/cmd/bootstrap"
	_ "goservicetemplate/pkg/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bootstrap.Start()

	// wait to ctrl + c
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGTERM)
	<-ch
}
