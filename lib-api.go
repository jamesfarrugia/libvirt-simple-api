package main

import "time"

// Meta app meta struct
type Meta struct {
	// time when started
	Started time.Time
}

// Config represents the service configuration
type Config struct {
	// listening host name
	Host string
	// listening port
	Port int
}

type Connection struct {
	ID  int
	URI string
}
