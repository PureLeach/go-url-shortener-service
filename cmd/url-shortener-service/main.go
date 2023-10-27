package main

import (
	"fmt"
	"url-shortener-service/cmd/internal/config"
)

func main() {
	cfg := config.ConfigLoad()
	fmt.Printf("cfg %#v\n", cfg)
}
