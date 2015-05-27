package arraycoder

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type ArrayEncoder []string

var seperator = ";;"

func SetSeperator(s string) {
	seperator = s
}

func (a ArrayEncoder) Encode() ([]byte, error) {
	l := len(a)
	if l == 0 {
		return nil, errors.New("Empty array")
	}

	var buffer bytes.Buffer
	for i := 0; i < l-1; i++ {
		buffer.WriteString(fmt.Sprintf("%s%s", a[i], seperator))
	}
	buffer.WriteString(fmt.Sprintf("%s", a[l-1]))

	return buffer.Bytes(), nil
}

func (a ArrayEncoder) Length() int {
	return len(a)
}

func Decode(b []byte) []string {
	s := string(b)
	return strings.Split(s, seperator)
}
