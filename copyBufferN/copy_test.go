package copy

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCopyBufferN(t *testing.T) {
	pattern := "abcdefghijklm1234567890"
	for i := 1; i < 1000000; i++ {
		src := bytes.NewBufferString(pattern)
		dest := bytes.NewBuffer(nil)
		_, err := CopyBufferN(dest, src, i)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(dest)
		if err != nil {
			t.Fatal(err)
		}
		by := []byte(pattern)
		if bytes.Compare(b, by) != 0 {
			t.Fatalf("%q != %q", b, by)
		}
	}
}
