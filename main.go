package main

import (
	"fmt"
	"log"
	"os"

	"github.com/L-chaCon/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	s := state{cfg: &cfg}
	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatalf("Not enought arguments %v", args)
		return
	}

	command := command{
		name:      args[1],
		arguments: args[2:],
	}

	err = commands.run(&s, command)
	if err != nil {
		log.Fatalf("Error running command %v", err)
	}
}
