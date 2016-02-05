package copy

import "io"

func CopyBufferN(dst io.Writer, src io.Reader, maxLen int) (written int64, err error) {
	if maxLen == 0 {
		panic("maxLen must be > 0")
	}
	buf := make([]byte, 32*1024)
	for {
		index := 0
		orig, er := src.Read(buf)
		nr, max := orig, orig
		for orig > 0 {
			if nr > maxLen {
				nr = index + maxLen
			} else {
				nr = max
			}
			nw, ew := dst.Write(buf[index:nr])
			if nw > 0 {
				written += int64(nw)
				orig -= nw
				index += nw
			}
			if ew != nil {
				err = ew
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}
