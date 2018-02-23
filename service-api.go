package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func doStartAPI(host string, port int) error {
	log.Info("Starting HTTP service at", host, ":", port, "...")
	htinfo := fmt.Sprintf("%s:%d", host, port)

	log.Info("Preparing API")
	router := doInitAPI()

	log.Info("Serving")
	err := http.ListenAndServe(htinfo, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return err
	}

	return nil
}

func doInitAPI() *httprouter.Router {
	router := httprouter.New()

	// meta
	router.GET("/", hxInfo)
	// config
	router.GET("/config", hxConfig)

	// connect to server
	router.POST("/connections", hxConnect)
	// get connections
	router.GET("/connections", hxConnections)

	// get domains
	router.GET("/connections/:connId/domains", hxDomains)

	return router
}

func hxInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	meta, err := json.Marshal(Meta{Started: app.started})

	if err != nil {
		log.Error("[API] - hxInfo - ", err)
	}

	fmt.Fprintf(w, "%s", meta)
}

func hxConfig(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	conf, err := json.Marshal(Config{
		Host: app.config.host,
		Port: app.config.port})

	if err != nil {
		log.Error("[API] - hxConfig - ", err)
	}

	fmt.Fprintf(w, "%s", conf)
}

func hxConnections(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	list := make([]Connection, 0, len(connections))
	for lvc := range connections {
		c := Connection{URI: lvc, ID: connections[lvc].id}
		list = append(list, c)
	}

	listJSON, err := json.Marshal(list)

	if err != nil {
		log.Error("[API] - hxConnections - ", err)
	}

	fmt.Fprintf(w, "%s", listJSON)
}

func hxConnect(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Error("Failed to close HTTP connection")
		}
	}()

	decoder := json.NewDecoder(r.Body)
	var rq rqConnection
	err := decoder.Decode(&rq)
	if err != nil {
		log.Error("Failed to understand request", err)
		w.WriteHeader(400)
		return
	}

	log.Info(rq)

	conn := lvConnect(rq.URL)
	c := Connection{URI: conn.url, ID: conn.id}
	connJSON, err := json.Marshal(c)

	if err != nil {
		log.Error("[API] - hxConnect - ", err)
	}

	fmt.Fprintf(w, "%s", connJSON)
}

func hxDomains(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("connId")
	conn := connections[id]
	domains, err := lvDomains(conn.conn)
	domainsJSON, err := json.Marshal(domains)

	if err != nil {
		log.Error("[API] - hxConnect - ", err)
	}

	fmt.Fprintf(w, "%s", domainsJSON)
}
