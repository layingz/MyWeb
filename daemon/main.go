package main

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/takama/daemon"
	"fmt"
	"strings"
	"superly.club/web/models"
)

type Service struct {
	Daemon daemon.Daemon
	Name string
}

func (s *Service) Manage() (string, error) {
	usage := fmt.Sprintf("Usage: %s install | remove | start | stop | status", s.Name)
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return s.Daemon.Install()
		case "remove":
			return s.Daemon.Remove()
		case "start":
			return s.Daemon.Start()
		case "stop":
			return s.Daemon.Stop()
		case "status":
			return s.Daemon.Status()
		default:
			return usage, nil
		}
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	killSignal := <-interrupt
	daemonlog.Println("Got signal:", killSignal)

	return "Service exited", nil
}

func newDeamon(name string, description string) daemon.Daemon {
	srv, err := daemon.New(name, description)
	if err != nil {
		daemonlog.Println("Error: ", err)
		os.Exit(1)
	}
	return srv
}

func NewDeamon(name string, description string) Service{
	return Service{Daemon: newDeamon(name ,description), Name: name}
}

var module []string
var daemonlog models.LogOp

func daemonParseModule(){
	if len(os.Args) > 2{
		m := os.Args[2]
		for _, date := range strings.Split(m, ","){
			module = append(module, date)
		}
	}
}

func main(){
	daemonlog = models.NewLog("daemonlog")
	defer daemonlog.Close()

	daemonParseModule()

	for _, name := range module{
		srv := NewDeamon(name, name + " Service")
		status, err := srv.Manage()
		if err != nil {
			daemonlog.Println(err)
		}
		fmt.Println(status)
	}
}
