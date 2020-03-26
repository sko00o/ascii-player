package player

import "time"

type timeBar struct {
	totalTime time.Duration

	width         int
	internalWidth int

	marker         string
	spacer         string
	leftDecorator  string
	rightDecorator string
}

func DefaultTimeBar(duration time.Duration) *timeBar {
	return &timeBar{
		totalTime: duration,

		width:         67,
		internalWidth: 65,

		marker:         "o",
		spacer:         "-",
		leftDecorator:  "<",
		rightDecorator: ">",
	}
}

func (t *timeBar) GetTimeBar(current time.Duration) string {
	pos := t.position(current)
	if pos >= t.internalWidth {
		pos = t.internalWidth - len(t.marker)
	}

	emptyBar := t.emptyBar()
	return emptyBar[:pos+len(t.leftDecorator)] +
		t.marker +
		emptyBar[pos+len(t.marker)+len(t.rightDecorator):]
}

func (t *timeBar) emptyBar() string {
	var barInternal string
	for i := 0; i < t.internalWidth; i++ {
		barInternal += t.spacer
	}
	return t.leftDecorator + barInternal + t.rightDecorator
}

func (t *timeBar) position(current time.Duration) int {
	return int(time.Duration(t.internalWidth) * current / t.totalTime)
}
