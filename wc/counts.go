package wc

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type Counts struct {
	Name  string `json:"name,omitempty"`
	Files int64  `json:"files,omitempty"`
	Lines int64  `json:"lines"`
	Words int64  `json:"words"`
	Bytes int64  `json:"bytes"`
}

func NewCounts(name string) *Counts {
	return &Counts{Name: name}
}

func (this *Counts) ReadFrom(r io.Reader) (n int64, err error) {
	initial := this.Bytes
	reader := bufio.NewReader(r)
	line := ""
	for err != io.EOF {
		line, err = reader.ReadString('\n')
		if err == nil || (err == io.EOF && len(line) > 0) {
			this.CountLine(line)
		}
	}
	if err == io.EOF {
		err = nil
	}
	return this.Bytes - initial, err
}
func (this *Counts) CountLine(line string) {
	if len(line) > 0 && line[len(line)-1] == '\n' {
		this.Lines++
	}
	this.Bytes += int64(len(line))
	line = strings.TrimSpace(line)
	if len(line) == 0 {
		return
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
			this.Words++
			inSpace = true
		}
	}
	if !inSpace {
		this.Words++
	}
}

func (this *Counts) Include(that *Counts) {
	this.Files++
	this.Lines += that.Lines
	this.Words += that.Words
	this.Bytes += that.Bytes
}
