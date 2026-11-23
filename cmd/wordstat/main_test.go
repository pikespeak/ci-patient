package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"testing/iotest"
)

func TestRun(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run(strings.NewReader("go go ci"), &out, &errOut); code != 0 || out.String() != "2\tgo\n1\tci\n" {
		t.Fatalf("code %d, out %q", code, out.String())
	}
}

func TestRunError(t *testing.T) {
	var out, errOut bytes.Buffer
	if code := run(iotest.ErrReader(errors.New("boom")), &out, &errOut); code != 1 || !strings.Contains(errOut.String(), "boom") {
		t.Fatalf("code %d, err %q", code, errOut.String())
	}
}
