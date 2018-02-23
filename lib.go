package main

import (
	"time"

	libvirt "github.com/libvirt/libvirt-go"
)

// Server configureation structure
type config struct {
	// listening host name
	host string
	// listening port
	port int
}

// program state
type appState struct {
	started time.Time
	config  *config
}

type rqConnection struct {
	URL string
}

type connection struct {
	url  string
	id   int
	conn *libvirt.Connect
}
