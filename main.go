package main

import (
	"fmt"
	"io"
	"os"
)

func GenerateScene() []string {
	return []string{}
}

func Run(stdout, stderr io.Writer, args []string) {
	for _, arg := range args {
		_, err := fmt.Fprintf(stdout, "%s", arg)
		if err != nil {
			_, err := fmt.Fprintf(stderr, "%v\n", err)
			if err != nil {
				os.Exit(1)
			}
		}
	}
}

func main() {
	Run(os.Stdout, os.Stderr, os.Args)
}
