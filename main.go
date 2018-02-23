package main

import (
	"os"
	"strconv"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

// Logger
var log = logging.MustGetLogger("libivrt-app")
var app appState

// Main function
func main() {
	app.started = time.Now()
	format := logging.MustStringFormatter(
		"%{color}%{time:15:04:05.000} %{shortfunc} ▶ \t%{level:.4s} %{id:03x}%{color:reset} %{message}",
	)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backend, backendFormatter)

	log.Info("libvirt API - James Farrugia 2018")
	conf := doInit(os.Args[1:])
	app.config = &conf

	lvInit()

	err := doStartAPI(conf.host, conf.port)
	if err != nil {
		log.Error(err.Error())
	}
}

// Goes through cmd args and sets up the host, port and service type
func doInit(args []string) config {
	conf := config{host: "127.0.0.1", port: 8080}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-host=") {
			conf.host = arg[6:]
		}

		if strings.HasPrefix(arg, "-port=") {
			portStr := arg[6:]
			var err error
			conf.port, err = strconv.Atoi(portStr)
			if err != nil {
				panic("Port must be a number")
			}
		}
	}

	return conf
}
