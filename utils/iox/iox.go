package iox

import (
	"io"
	"os"
)

func FileCopy(dst, src string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()
	dstF, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer dstF.Close()

	_, err = io.Copy(dstF, srcF)
	return err
}

func ReadToFile(dst string, src io.Reader) error {
	fd, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fd, src)
	return err
}

func WriteFromFile(dst io.Writer, src string) error {
	fd, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = io.Copy(dst, fd)
	return err
}
