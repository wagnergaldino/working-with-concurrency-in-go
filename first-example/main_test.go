package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_testSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("test something", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "test something") {
		t.Errorf("Expected output: test something - isn't available")
	}

}
