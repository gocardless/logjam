package main

import (
	"bytes"
	"net"
	"testing"
)

type StubConn struct {
	net.Conn

	buffer *bytes.Buffer
}

func (sc StubConn) Write(p []byte) (int, error) {
	return sc.buffer.Write(p)
}

func TestWriteWithBackoff(t *testing.T) {
	conn := StubConn{buffer: new(bytes.Buffer)}
	s := Shipper{conn}

	s.WriteWithBackoff([]byte("hello"), 125)

	if bytes.Compare(conn.buffer.Bytes(), []byte("hello")) != 0 {
		t.Fatal("Write Mismatch")
	}
}
