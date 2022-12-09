package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransformFunc func(string) string

type Server struct {
	filenameTransformFunc TransformFunc
}

func (s *Server) handleRequest(filename string) error {

	newFilename := s.filenameTransformFunc(filename)
	fmt.Println("new file name is ", newFilename)
	return nil
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func GGPrefixFilename(prefix string) TransformFunc {
	return func(filename string) string {
		return prefix + filename
	}
}
func main() {

	s := &Server{
		filenameTransformFunc: GGPrefixFilename("Bob_"),
	}
	s.handleRequest("cool_picture.jpg")
}
