package system

import(
	"bufio"
	"io"
	"os"

	"bytes"
	"strings"
)

func ReadTemplate(file string) *string {

	fi, err := os.Open(strings.Join([]string{"templates/", file}, ""))
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(fi)
	buf := make([]byte, 4096)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}

	n := bytes.IndexByte(buf, 0)
	str := string(buf[:n])

	return &str
}