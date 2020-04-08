package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/ptime/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewPTimeCommand()

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
