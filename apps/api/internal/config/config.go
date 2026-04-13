package config

import (
)

type Config struct {
	DBPort     string
	DBUrl	   string
	ENV		   string
}

var cfg *Config

func load() {
	
}