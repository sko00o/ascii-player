package video

import (
	"testing"
)

func TestPadding(tst *testing.T) {
	tks := []string{
		"1",
		"12",
		"123",
		"1234",
	}

	tst.Log("padLeft")
	for _, t := range tks {
		out := padLeft(t, 5)
		tst.Log(out)
	}

	tst.Log("padRight")
	for _, t := range tks {
		out := padRight(t, 5)
		tst.Log(out)
	}
}
