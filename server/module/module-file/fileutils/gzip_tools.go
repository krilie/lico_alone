package fileutils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func MustGzipBytes(in []byte) []byte {
	if out, err := GzipBytes(in); err != nil {
		panic(err)
	} else {
		return out
	}
}
func MustUnGzipBytes(in []byte) []byte {
	if out, err := UnGzipBytes(in); err != nil {
		panic(err)
	} else {
		return out
	}
}

func GzipBytes(in []byte) (out []byte, err error) {
	var outBuffer bytes.Buffer
	w := gzip.NewWriter(&outBuffer)
	_, err = w.Write(in)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	return outBuffer.Bytes(), nil
}

func UnGzipBytes(in []byte) (out []byte, err error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return all, nil
}
