package main

import (
	cli "github.com/brunobach/cli/cmd"
)

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
