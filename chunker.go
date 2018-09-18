package chunker

func main() {}

func ToChunks(toBeChunked string, noOfBytes int) [][]byte {
	var chunks [][]byte

	f := []byte(toBeChunked)
	prev := 0
	times := float64(len(f)) / float64(noOfBytes)

	if times != float64(int(times)) {
		times = float64(int(times) + 1)
	}

	for i := 0; i < int(times); i++ {
		b := f[prev : prev+noOfBytes]
		chunks = append(chunks, []byte(string(b)))
		prev += noOfBytes
	}
	return chunks
}

func FromChunks(toBeUnchunked [][]byte) string {
	var unChunked []byte

	for i := 0; i < len(toBeUnchunked); i++ {
		for j := 0; j < len(toBeUnchunked[i]); j++ {
			unChunked = append(unChunked, toBeUnchunked[i][j])
		}
	}

	return string(RemoveTrailingZeros(unChunked))
}

func RemoveTrailingZeros(arr []byte) []byte {
	zeros := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeros++
		}
	}

	return arr[:len(arr)-zeros]
}
