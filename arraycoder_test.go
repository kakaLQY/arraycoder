package arraycoder

import "testing"

func TestAll(t *testing.T) {
	s := ArrayEncoder{"hello", "world", "你好"}
	b, err := s.Encode()
	if err != nil {
		t.Fatal(err)
	}

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
