package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
	chunks := toChunks()
	fromChunks(chunks)
}

func toChunks() [][]byte {
	var chunks [][]byte
	noOfBytes := 4

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	l, _ := f.Stat()
	size := int(l.Size())
	r := bufio.NewReader(f)
	b := make([]byte, noOfBytes)

	for i := 0; i < ((size/noOfBytes) + 1) ; i++ {
		num, err := r.Read(b)
		if num != noOfBytes && num != 0 {
			b = b[0:num-1]
		}

		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		chunks = append(chunks, []byte(string(b)))
	}
	return chunks
}

func fromChunks(b [][]byte) string {
	var unChunked []byte

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			unChunked = append(unChunked, b[i][j])
		}
	}
	return string(unChunked)
}
