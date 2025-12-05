package main

import (
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA256Sig("02-Data_Structures_and_REST_APIs/07-Composing_io_reader_and_writer/sha256/http.log.gz"))
	fmt.Println(SHA256Sig("02-Data_Structures_and_REST_APIs/07-Composing_io_reader_and_writer/sha256/sha256.go"))
}

// SHA256Sig returns SHA256 signature of uncompressed file
// Exercise: Decompress only if file name ends with ".gz"
// cat http.log.gz | gunzip | sha256sum
func SHA256Sig(fileName string) (string, error) {

	// cat http.log.gz
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		// | gunzip
		// BUG: Creates new "r" that is only in "if" scope
		// shadowing
		// r, err := gzip.NewReader(file)
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("%q - gzip: %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	// Decompress only if file name ends with ".gz"
	/*
		var ext = filepath.Ext("fileName")
		if ext != ".gz" {
			return "Wrong file extension", nil
		} else {
			var f []byte
			if _, err := io.Copy(f, r); err != nil {
				return "", fmt.Errorf("%q - copy: %w", fileName, err)
			}
			err := os.WriteFile("http.log", f, 0664)
			if err != nil {
				return "", err
			}
		}
	*/

	// | sha256sum
	w := sha256.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

/*
Go
type Reader interface {
	Read(p []byte) (n int, err error)
}

Python
type Reader interface {
	Read(n int) ([]byte, err error)
}
Better performance in Go then Python
*/
