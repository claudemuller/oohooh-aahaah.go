package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/claudemuller/oohooh-aahaah-go/internal/tools/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("error retrieving system user: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", usr.Username)
	fmt.Println("Feel free to type in commands.")

	repl.Start(os.Stdin, os.Stdout)
}
