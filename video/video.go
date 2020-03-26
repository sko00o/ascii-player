package video

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	lineSep = "\n"
)

type Frame struct {
	Content     []string
	DisplayTime time.Duration
}

type Video struct {
	Frames    []Frame
	TotalTime time.Duration

	width  int
	height int

	loaded bool
}

func DefaultVideo() *Video {
	return &Video{
		width:  67,
		height: 13,

		Frames: []Frame{
			{
				Content:     []string{"No video yet."},
				DisplayTime: 1,
			},
		},
	}
}

func (v *Video) Load(filePath string) error {
	if v.loaded {
		return nil
	}

	// must have enough ram for file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to load file %s: %v\n", filePath, err)
	}

	frames, err := v.extract(data)
	if err != nil {
		return fmt.Errorf("failed to extract: %v\n", err)
	}

	v.Frames = frames
	return nil
}

func (v *Video) extract(data []byte) ([]Frame, error) {
	var frames []Frame
	lines := strings.Split(string(data), lineSep)

	// one line for display time
	frameHeight := v.height + 1
	for i := 0; i+frameHeight <= len(lines); i += frameHeight {
		// parse display time
		dStr := strings.TrimSpace(lines[i])
		dInt, err := strconv.ParseInt(dStr, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("parse frame duration error: %v", err)
		}
		frameDuration := time.Duration(dInt)
		v.TotalTime += frameDuration

		content := lines[i+1 : i+frameHeight]
		for i := range content {
			content[i] = strings.TrimRightFunc(content[i], unicode.IsSpace)
			content[i] = padRight(content[i], v.width)
		}

		frames = append(frames, Frame{
			Content:     content,
			DisplayTime: frameDuration,
		})
	}
	return frames, nil
}

func padLeft(s string, width int) string {
	return fmt.Sprintf(fmt.Sprintf("%%%ds", width), s)
}

func padRight(s string, width int) string {
	return fmt.Sprintf(fmt.Sprintf("%%-%ds", width), s)
}
