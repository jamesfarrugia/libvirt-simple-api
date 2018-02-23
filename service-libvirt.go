package main

import (
	"time"

	"github.com/libvirt/libvirt-go"
)

var connections map[string]*connection

func lvInit() {
	connections = make(map[string]*connection)
}

func lvConnect(url string) *connection {
	conn := connections[url]
	if conn != nil {
		log.Info("Connection already established")
		return conn
	}
	connect, err := libvirt.NewConnect(url)
	if err != nil {
		log.Error("Failed to connect", err)
	}
	newConn := connection{url: url, conn: connect, id: int(time.Now().Unix())}
	connections[url] = &newConn
	return &newConn
}

func lvDomains(conn *libvirt.Connect) ([]libvirt.Domain, error) {
	return conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
}
