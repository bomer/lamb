package main

import "testing"

func TestHash(t *testing.T) {
	testHash := CreateKey("James")
	if testHash != "9345a35a6fdf174dff7219282a3ae4879790dbb785c70f6fff91e32fafd66eab" {
		t.Error(testHash)
		t.Error("Hash didn't works")
	}

}
