package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
)

func GenerateScene(data Words) []string {
	intVerb := rand.Intn(100)
	intAdj := rand.Intn(100)
	intNoun := rand.Intn(100)
	return []string{data.Nouns[intNoun], data.Adjectives[intAdj], data.Verbs[intVerb]}
}

func setChaos() int {
	return 10000
}

func Run(stdout, stderr io.Writer, args []string, data Words) {
	sceneWords := GenerateScene(data)
	_, err := fmt.Fprintf(stdout, "\n=====ORC======\nScene:\n %s, %s, %s \n\n", sceneWords[0], sceneWords[1], sceneWords[2])
	if err != nil {
		_, errStderr := fmt.Fprintf(stderr, "%v\n", err)
		if errStderr != nil {
			os.Exit(1)
		}
	}
}

func main() {
	Run(os.Stdout, os.Stderr, os.Args[1:], Data)
}
