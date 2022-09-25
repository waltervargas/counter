package counter_test

import (
	"bytes"
	"testing"

	"github.com/waltervargas/counter"
)

func TestLines(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := counter.NewCounter(
		counter.WithInput(inputBuf),
		counter.WithOutput(inputBuf),
	)
	if err != nil {
		panic("internal error")
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := counter.NewCounter(
		counter.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgsEmpty(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := counter.NewCounter(
		counter.WithInput(inputBuf),
		counter.FromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWords(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2 words\n3 this time")
	c, err := counter.NewCounter(
		counter.WithInput(inputBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", got, want)
	}
}

func TestWordCount(t *testing.T) {
	args := []string{"-w", "testdata/three_lines.txt"}
	c, err := counter.NewCounter(
		counter.FromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
