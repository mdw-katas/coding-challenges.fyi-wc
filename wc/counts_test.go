package wc

import (
	"errors"
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
	stats Stats
}

func (this *Suite) process(text string) (result *Stats) {
	result = NewStats("")
	n, err := result.ReadFrom(strings.NewReader(text))
	this.So(n, should.Equal, len(text))
	this.So(err, should.BeNil)
	return result
}
func (this *Suite) TestEmpty() {
	this.So(this.process(""), should.Equal, &Stats{})
}
func (this *Suite) TestSpace() {
	this.So(this.process(" "), should.Equal, &Stats{Bytes: 1})
}
func (this *Suite) TestSingleLineMultipleWords() {
	this.So(this.process("1  22\t333\v4444\f55555 "), should.Equal, &Stats{
		Bytes: 21,
		Words: 5,
	})
}
func (this *Suite) TestMultipleLinesAndWords() {
	result := this.process(strings.Join([]string{
		"1",
		"2 2",
		"3 3 3",
		"4 4 4 4",
		"5 5 5 5 5",
	}, "\n"))
	this.So(result, should.Equal, &Stats{
		Bytes: 29,
		Words: 15,
		Lines: 4,
	})
}
func (this *Suite) TestInclude() {
	a := NewStats("a")
	b := NewStats("b")
	a.Bytes, a.Words, a.Lines = 5, 4, 3
	b.Bytes, b.Words, b.Lines = 5, 4, 3

	a.Include(b)

	this.So(a, should.Equal, &Stats{
		Name:  "a",
		Bytes: 5 + 5,
		Words: 4 + 4,
		Lines: 3 + 3,
		Files: 1,
	})
}
func (this *Suite) TestReadError() {
	boink := errors.New("boink")
	result := NewStats("")
	n, err := result.ReadFrom(NewErringReader("hi", boink))
	this.So(n, should.Equal, 2)
	this.So(err, should.Equal, boink)
}

type ErringReader struct {
	content string
	err     error
}

func (this *ErringReader) Read(p []byte) (n int, err error) {
	return copy(p, this.content), this.err
}

func NewErringReader(content string, err error) *ErringReader {
	return &ErringReader{content: content, err: err}
}
