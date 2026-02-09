package main

import (
	"bytes"
	"fmt"
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
		Run(&stdout, &stderr, args)
		if len(stdout.String()) == 0 {
			t.Fatalf("expected len greater than 0 printed")
		}
	})

	t.Run("Run errors and prints to stderr if stdout fails", func(t *testing.T) {
		stdoutErr := errStdout{}
		stderr := bytes.Buffer{}
		args := []string{"program", "one", "two"}
		Run(&stdoutErr, &stderr, args)
		if len(stderr.String()) == 0 {
			t.Fatalf("expected len greater than 0 printed")
		}
	})

	t.Run("Generate returns a noun, a verb and adjective as a []string len == 3", func(t *testing.T) {
		got := GenerateScene()
		want := 3
		if len(got) != want {
			t.Fatalf("got %d want %d", len(got), want)
		}
	})
}
