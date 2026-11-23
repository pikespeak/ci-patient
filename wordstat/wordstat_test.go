package wordstat

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"testing/iotest"
)

func TestTop(t *testing.T) {
	got, err := Top(strings.NewReader("Build, test, build!\nÜber test a"), 3)
	want := []Entry{{"build", 2}, {"test", 2}, {"a", 1}}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, %v; want %v", got, err, want)
	}
}

func TestTopEmpty(t *testing.T) {
	got, err := Top(strings.NewReader(""), 5)
	if err != nil || len(got) != 0 {
		t.Fatalf("got %v, %v", got, err)
	}
}

func TestTopRejectsBadN(t *testing.T) {
	if _, err := Top(strings.NewReader("a"), 0); err == nil {
		t.Fatal("expected an error")
	}
}

func TestTopReadError(t *testing.T) {
	boom := errors.New("boom")
	if _, err := Top(iotest.ErrReader(boom), 1); !errors.Is(err, boom) {
		t.Fatalf("got %v", err)
	}
}
