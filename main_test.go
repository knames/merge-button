package main

import (
	"testing"
)

func TestDelete(t *testing.T)  {
	title1 := "Issue 3423: A dope PR"
	title2 := "Issue 3950: Another"

	ts := []*string{&title1, &title2}
	titles = ts

	delete(0)

	if len(titles) != 1 {
		t.Error("Expected to be length 1.")
	}
}