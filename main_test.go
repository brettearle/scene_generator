package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

type errStdout struct{}

func (e *errStdout) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failed to stdout")
}

func TestMain(t *testing.T) {
	t.Run("Run prints to stdout", func(t *testing.T) {
		stdout := bytes.Buffer{}
		stderr := bytes.Buffer{}
		args := []string{"program", "one", "two"}
		Run(&stdout, &stderr, args, Data)
		if len(stdout.String()) == 0 {
			t.Fatalf("expected len greater than 0 printed")
		}
	})

	t.Run("Run errors and prints to stderr if stdout fails", func(t *testing.T) {
		stdoutErr := errStdout{}
		stderr := bytes.Buffer{}
		args := []string{"program", "one", "two"}
		Run(&stdoutErr, &stderr, args, Data)
		if len(stderr.String()) == 0 {
			t.Fatalf("expected len greater than 0 printed")
		}
	})

	t.Run("Generate returns a noun, a adjective and verb as a []string len(3)", func(t *testing.T) {
		got := GenerateScene(Data)
		want := 3
		if len(got) != want {
			t.Fatalf("got %d want %d", len(got), want)
		}
	})

	t.Run("Generate returns a noun, a adjective and verb", func(t *testing.T) {
		got := GenerateScene(Data)

		nouns := strings.Join(Data.Nouns, ",")
		adjectives := strings.Join(Data.Adjectives, ",")
		verbs := strings.Join(Data.Verbs, ",")

		fmt.Println(got)
		if !strings.Contains(nouns, got[0]) {
			t.Errorf("got %s which is not in nouns", got[0])
		}
		if !strings.Contains(adjectives, got[1]) {
			t.Errorf("got %s which is not in adjectives", got[1])
		}
		if !strings.Contains(verbs, got[2]) {
			t.Errorf("got %s which is not in verbs", got[2])
		}
	})

	t.Run("Set chaos returns a number between 1-10", func(t *testing.T) {
		got := setChaos()
		if got < 0 || got > 10 {
			t.Fatalf("got %d want between 1 and 10", got)
		}
	})
}
