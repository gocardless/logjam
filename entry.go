package main

import (
	"encoding/json"
	"os"
)

var Hostname string = "unknown"

type Entry struct {
	Host    string `json:"host"`
	Message string `json:"message"`
}

func (e *Entry) ToJSON() []byte {
	e.Host = Hostname
	dump, _ := json.Marshal(e)
	return dump
}

func init() {
	Hostname, _ = os.Hostname()
}
