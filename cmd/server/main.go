package main

import (
	"flag"
	"fmt"
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
		WriteTimeout: 15,
		ReadTimeout:  15,
		DatabaseRoot: "database",
	}

	srv, err := NewServer(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	srv.start()
}
