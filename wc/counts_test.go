package wc

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
	stats Counts
}

func (this *Suite) process(text string) (result *Counts) {
	result = NewCounts("")
	n, err := result.ReadFrom(strings.NewReader(text))
	this.So(n, should.Equal, len(text))
	this.So(err, should.BeNil)
	return result
}
func (this *Suite) TestEmpty() {
	this.So(this.process(""), should.Equal, &Counts{})
}
func (this *Suite) TestSpace() {
	this.So(this.process(" "), should.Equal, &Counts{Bytes: 1})
}
func (this *Suite) TestSingleLineMultipleWords() {
	this.So(this.process("1 22\t333\v4444\f55555 "), should.Equal, &Counts{
		Bytes: 20,
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
	this.So(result, should.Equal, &Counts{
		Bytes: 29,
		Words: 15,
		Lines: 4,
	})
}
