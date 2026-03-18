package main

import "fmt"

type AppConfig struct {
	Host  string
	Port  int
	Debug bool
}

func main() {
	var cfg AppConfig
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("Host: %v\n", cfg.Host)
	fmt.Printf("Port: %v\n", cfg.Port)
	fmt.Printf("Debug: %v\n", cfg.Debug)

	if cfg.Host == "" {
		fmt.Println("Host is empty")
	}
	if cfg.Port == 0 {
		fmt.Println("Port is empty")
	}
	if cfg.Debug == false {
		fmt.Println("Debug is empty")
	}
}
