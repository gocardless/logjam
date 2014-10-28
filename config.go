package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Bind           string   `json:"bind"`        // interface to bind to (0.0.0.0 for all)
	Port           int      `json:"port"`        // listen port
	Server         string   `json:"server"`      // remove server to publish logs to
	DiskBufferPath string   `json:"buffer"`      // path for disk buffer
	BufferSize     int      `json:"buffer_size"` // queue length of memory bufer
	TruncatePeriod int      `json:"truncate"`    // cleanup buffer every x seconds
	Files          []string `json:"files"`       // files to include in publish
}

// read config file
func ReadConfigFile(path string) (cfg Config, err error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = json.Unmarshal(file, &cfg)
	return
}
