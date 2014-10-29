package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/ActiveState/tail"
)

type Shipper struct {
	net.Conn
}

// create a new shipper client
func NewShipper(proto string, addr string) (*Shipper, error) {
	conn, err := net.Dial(proto, addr)
	if err != nil {
		return nil, err
	}

	return &Shipper{conn}, nil
}

// write to socket with exponential backoff in milliseconds
func (s *Shipper) WriteWithBackoff(p []byte, initial int) {
	var timeout time.Duration = time.Duration(initial) * time.Millisecond

	for {
		_, err := s.Write(p)
		if err != nil {
			timeout = timeout * 2
			time.Sleep(timeout)
			continue
		}

		return
	}
}

// ship entries to remote log server
func (s *Shipper) Ship(filename string) {
	t, err := tail.TailFile(filename, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		log.Printf("Shipper: Error: %s\n", err)
		return
	}

	for line := range t.Lines {
		s.WriteWithBackoff([]byte(line.Text), 125)
	}
}

// truncate a file every period
func (s *Shipper) TruncateEvery(filename string, period time.Duration) {
	for {
		time.Sleep(period)

		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("Shipper: Truncate: Error: %s\n", err)
			continue
		}
		file.Close()

		log.Printf("Shipper: Truncate: %s\n", filename)
	}
}
