package main

import "time"

//ServerConfig is our config
type ServerConfig struct {
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}
