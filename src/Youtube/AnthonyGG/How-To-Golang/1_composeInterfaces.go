package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type HashReader interface {
	io.Reader
	hash() string
}

func main() {

	payload := []byte("High value soft. engineers")
	// hashAndBroadcast(bytes.NewReader(payload)) //instead of this..
	hashAndBroadcast(NewHashReader(payload))
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r HashReader) error {

	hash := r.hash()
	fmt.Println(hash)
	return broadcast(r)
}

// func hashAndBroadcast(r io.Reader) error {

// 	b, err := ioutil.ReadAll(r)

// 	if err != nil {
// 		return err
// 	}
// 	hash := sha1.Sum(b)
// 	fmt.Println(hex.EncodeToString(hash[:]))

// 	// you are passing r(reader) which has already read the content. you won't see the string bytes size
// 	return broadcast(r)
// }

func broadcast(r io.Reader) error {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("strings of bytes: ", string(b))
	return nil
}
