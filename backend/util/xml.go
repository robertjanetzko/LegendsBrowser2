package util

import (
	"bufio"
)

type XMLParser struct {
	reader      *bufio.Reader
	scratch     *scratch
	selfClose   bool
	lastElement string
}

func NewXMLParser(r *bufio.Reader) *XMLParser {
	x := &XMLParser{
		reader:  r,
		scratch: &scratch{data: make([]byte, 1024)},
	}
	x.skipDeclerations()
	return x
}

func (x *XMLParser) skipDeclerations() error {
	for {
		b, err := x.reader.ReadByte()
		if err != nil {
			return err
		}
		if b == '>' {
			return nil
		}
	}
}

type TokenType int

const (
	StartElement TokenType = iota
	EndElement
)

func (x *XMLParser) Token() (TokenType, string, error) {
	if x.selfClose {
		x.selfClose = false
		return EndElement, x.lastElement, nil
	}

	var (
		f, c bool
	)
	for {
		b, err := x.reader.ReadByte()
		if err != nil {
			return 0, "", err
		}
		if b == '<' {
			f = true
			x.scratch.reset()

			b, err := x.reader.ReadByte()
			if err != nil {
				return 0, "", err
			}

			if b == '/' {
				c = true
			} else {
				x.scratch.add(b)
			}

		} else if b == '>' {
			bs := x.scratch.bytes()
			if bs[len(bs)-1] == '/' {
				x.selfClose = true
				x.lastElement = string(bs[:len(bs)-1])
				return StartElement, x.lastElement, nil
			} else {
				if c {
					return EndElement, string(bs), nil
				} else {
					return StartElement, string(bs), nil
				}
			}

		} else if f {
			x.scratch.add(b)
		}
	}
}

func (x *XMLParser) Value() (string, error) {
	if x.selfClose {
		x.selfClose = false
		return "", nil
	}

	x.scratch.reset()
	var (
		f bool
	)
	for {
		b, err := x.reader.ReadByte()
		if err != nil {
			return "", err
		}
		if b == '<' {
			f = true
		} else if f && b == '>' {
			return string(x.scratch.bytes()), nil

		} else if !f {
			x.scratch.add(b)
		}
	}
}

func (x *XMLParser) Skip() error {
	depth := 0
	for {
		t, _, err := x.Token()
		if err != nil {
			return err
		}
		switch t {
		case StartElement:
			depth++
		case EndElement:
			if depth == 0 {
				return nil
			}
			depth--
		}
	}
}

// scratch taken from
//https://github.com/bcicen/jstream
type scratch struct {
	data []byte
	fill int
}

// reset scratch buffer
func (s *scratch) reset() { s.fill = 0 }

// bytes returns the written contents of scratch buffer
func (s *scratch) bytes() []byte { return s.data[0:s.fill] }

// grow scratch buffer
func (s *scratch) grow() {
	ndata := make([]byte, cap(s.data)*2)
	copy(ndata, s.data[:])
	s.data = ndata
}

// append single byte to scratch buffer
func (s *scratch) add(c byte) {
	if s.fill+1 >= cap(s.data) {
		s.grow()
	}

	s.data[s.fill] = c
	s.fill++
}
