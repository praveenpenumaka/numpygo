package loaders

import (
	"io"
)

type NPYHeader struct {
}

type NPY struct {
	MAGIC        [6]byte
	MAJORVERSION byte
	MINORVERSION byte
	HEADERLEN    uint16
	Header       NPYHeader
}

func LoadNPY(reader io.Reader) error {
	return nil
}

func SaveNPY(write io.Reader) {

}
