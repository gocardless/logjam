package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"regexp"

	"github.com/ActiveState/tail"
)

var (
	ValidJSON = regexp.MustCompile("\\{.*\\}")
)

type Receiver struct {
	Host string // listen address
	Port int    // listen port

	messages chan []byte // incomming messages
}

// create a new receiver server
func NewReceiver(host string, port, bufferSize int) Receiver {
	return Receiver{
		Host: host,
		Port: port,

		messages: make(chan []byte, bufferSize),
	}
}

func (r *Receiver) ListenAndServe() error {
	addr := net.UDPAddr{Port: r.Port, IP: net.ParseIP(r.Host)}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		b := scanner.Bytes()

		if !ValidJSON.Match(b) {
			log.Printf("Receiver: Error: Invalid Message\n")
			continue
		}

		r.messages <- b
	}

	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (r *Receiver) ListenToFile(path string) error {
	t, err := tail.TailFile(path, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		return err
	}

	for line := range t.Lines {
		m := Entry{Message: line.Text}
		r.messages <- m.ToJSON()
	}

	return nil
}

// write entries on messages channel to filename
func (r *Receiver) WriteToFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Writer: Error: %s\n", err)
		return
	}
	defer file.Close()

	for {
		file.Write(<-r.messages)
		file.Write([]byte("\n"))
	}
}
