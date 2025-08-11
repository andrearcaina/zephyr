package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/andrearcaina/zephyr/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is Zephyr, a W.I.P programming language.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
