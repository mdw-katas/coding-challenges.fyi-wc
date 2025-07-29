package wc

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type Stats struct {
	Name  string `json:"name,omitempty"`
	Files int64  `json:"files,omitempty"`
	Lines int64  `json:"lines"`
	Words int64  `json:"words"`
	Bytes int64  `json:"bytes"`
}

func NewStats(name string) *Stats {
	return &Stats{Name: name}
}

func (this *Stats) ReadFrom(r io.Reader) (n int64, err error) {
	reader := bufio.NewReader(r)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if len(line) > 0 {
			lines, words, bytes := this.processLine(line)
			this.Lines += lines
			this.Words += words
			this.Bytes += bytes
			n += bytes
		}
		if err == io.EOF {
			return n, nil
		}
		if err != nil {
			return n, err
		}
	}
}
func (this *Stats) processLine(line string) (lines, words, bytes int64) {
	if len(line) > 0 && line[len(line)-1] == '\n' {
		lines++
	}
	bytes += int64(len(line))
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return lines, words, bytes
	}
	inSpace := unicode.IsSpace(rune(line[0]))
	for x := 1; x < len(line); x++ {
		isSpace := unicode.IsSpace(rune(line[x]))
		if inSpace && isSpace {
			continue
		}
		if inSpace {
			inSpace = false
			continue
		}
		if isSpace {
			words++
			inSpace = true
		}
	}
	if !inSpace {
		words++
	}
	return lines, words, bytes
}

func (this *Stats) Include(that *Stats) {
	this.Files++
	this.Lines += that.Lines
	this.Words += that.Words
	this.Bytes += that.Bytes
}
