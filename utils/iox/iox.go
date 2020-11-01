package iox

import (
	"io"
	"os"
)

func CopyToFile(src io.Reader, dist string) error {
	fd, err := os.OpenFile(dist, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fd, src)
	return err
}

func CopyFromFile(src string, dist io.Writer) error {
	fd, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(dist, fd)
	return err
}
