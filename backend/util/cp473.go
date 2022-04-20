package util

import "io"

var cp437 = []byte("         \t\n  \r                   !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ CueaaaaceeeiiiAAEaAooouuyOU    faiounN                                                                                                ")

// var cp437 = []byte("         \t\n  \r                   !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~ ÇüéâäàåçêëèïîìÄÅÉæÆôöòûùÿÖÜ¢£¥₧ƒáíóúñÑ                                                                                                ")

type ConvertReader struct {
	r    io.Reader
	read int
}

func NewConvertReader(r io.Reader) *ConvertReader {
	return &ConvertReader{r: r}
}

func (c *ConvertReader) Read(b []byte) (n int, err error) {
	n, err = c.r.Read(b)
	if c.read == 0 && n > 35 {
		copy(b[30:35], []byte("UTF-8"))
	}
	c.read += n
	for i := range b {
		b[i] = cp437[b[i]]
	}
	return n, err
}

func ConvertCp473(b []byte) string {
	for i := range b {
		b[i] = cp437[b[i]]
	}
	return string(b)
}
