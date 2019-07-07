package tail

import (
	"testing"
)

func TestTail(t *testing.T) {
	s := Tail(3, "./forTest.txt")
	if s[0] != "This is line-18" {
		t.Fatal("s[0] is not correct")
	}
	if s[1] != "This is line-19" {
		t.Fatal("s[1] is not correct")
	}
	if s[2] != "This is line-20" {
		t.Fatal("s[2] is not correct")
	}
}
