package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/nsf/termbox-go"
)

var (
	row int
	col int
)

func echo(str []byte) {
	for _, v := range str {
		if v == 0 {
			break
		}
		if v == 13 {
			row++
			col = 0
			break
		}
		termbox.SetCell(col, row, rune(v), termbox.ColorWhite, termbox.ColorBlack)
		col++
	}
	termbox.Flush()
}
func main() {
	if err := termbox.Init(); err != nil {
		logrus.Fatal(err)
	}
	defer termbox.Close()

	for {
		buff := make([]byte, 32)

		termbox.PollRawEvent(buff)
		echo(buff)
	}

}
