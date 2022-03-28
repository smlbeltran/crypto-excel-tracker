package main

import (
	"testing"
	"testing/fstest"
)

func TestReadFileDir(t *testing.T) {

	fs := fstest.MapFS{
		"crypto.txt": {Data: []byte(`
		bitcoin
		eos
		xrp`)},
	}

	got := readCoinsFromFile(fs)

	want := []string{"bitcoin", "eos", "xrp"}

	if len(got) != len(want) {
		t.Errorf("want: %q, got:%q", len(want), len(got))
	}

}
