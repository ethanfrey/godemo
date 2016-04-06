package main

import (
//  "fmt"
  "os"
// "strings"
  "io"
)

type MyReader struct{
    size int
}


func (r *MyReader)Read(out []byte) (int, error) {
    // read as much as we can before closing it
    n := len(out)
    if r.size < n {
        n = r.size
    }
    r.size -= n

    for i := 0; i < n; i++ {
      out[i] = 'A' + byte(i)
    }

    if n == 0 {
        return -1, io.EOF
    } else {
        return n, nil
    }
}


type Rot13Reader struct {
    r io.Reader
}

func Rot13(ch byte) byte {
    for _, offset := range [2]byte{64, 96} {
        if offset < ch && ch < offset+27 {
            return (ch - offset + 13) % 26 + offset
        }
    }
    return ch
}

func (r *Rot13Reader) Read(out []byte) (int, error) {
    n, err := r.r.Read(out)
    for i := 0; i < n; i++ {
        out[i] = Rot13(out[i])
    }
    return n, err
}


func main() {
    data := MyReader{30}
    // data := strings.NewReader("Hello, World!")
    filter := Rot13Reader{&data}
    // original := Rot13Reader{&filter}

    io.Copy(os.Stdout, &filter)
    // io.Copy(os.Stdout, &original)
}
