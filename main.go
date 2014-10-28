package main

import (
	"flag"
	"log"
	"time"
)

var (
	ConfigPath = flag.String("config", "/etc/log-stream.json", "Config File Path")
)

func main() {
	flag.Parse()

	config, err := ReadConfigFile(*ConfigPath)
	if err != nil {
		log.Fatalf("Config: Error: %s\n", err)
	}

	r := NewReceiver(config.Bind, config.Port, config.BufferSize)
	s := NewShipper(config.Server)

	err = s.Dial()
	if err != nil {
		log.Fatalf("Shipper: Error: %s\n", err)
	}
	log.Printf("Shipper: Connected: %s\n", config.Server)

	go r.WriteToFile(config.DiskBufferPath)
	go s.Ship(config.DiskBufferPath)
	go s.TruncateEvery(config.DiskBufferPath, time.Duration(config.TruncatePeriod)*time.Second)

	// Ship Files
	for _, path := range config.Files {
		go r.ListenToFile(path)
	}

	// Ship Socket
	err = r.ListenAndServe()
	if err != nil {
		log.Fatalf("Receiver: Error: %s\n", err)
	}
}
