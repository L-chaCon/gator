package main

import "fmt"
import "github.com/L-chaCon/gator/internal/config"

func main() {
	fmt.Println("Reading config...")
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config %v", err)
	}

	cfg.SetUser("chaCon")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config %v", err)
	}
	fmt.Println(cfg)
}
