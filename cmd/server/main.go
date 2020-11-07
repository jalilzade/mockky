package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	ip := flag.String("ip", "127.0.0.1", "ip that will server host")
	port := flag.Int("port", 8080, "port that will server listen")
	helpflag := flag.Bool("help", false, "shows help")
	flag.Parse()

	if *helpflag {
		flag.Usage()
	}

	config := &ServerConfig{
		Address:      fmt.Sprintf("%s:%d", *ip, *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv := NewServer(config)
	srv.start()
}
