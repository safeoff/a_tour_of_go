package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error){
	n, e := r.r.Read(b)
	for i, c := range b{
		if c >= 'A' && c <= 'Z' {
			b[i] = (c - 'A' + 13) % 26 + 'A'
		} else if c >= 'a' && c <= 'z'{
			b[i] = (c - 'a' + 13) % 26 + 'a'
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

