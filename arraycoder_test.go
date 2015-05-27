package arraycoder

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	SetSeperator("||")
	s := ArrayEncoder{"hello", "world", "你好"}
	b, err := s.Encode()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("result:", string(b))
	arr := Decode(b)
	if len(arr) != len(s) {
		t.Fatal("Decode is failed.")
	}
	for i, v := range arr {
		if v != s[i] {
			t.Fatal("Decode is failed.")
		}
	}
}
