package wc

import (
	"io/fs"
)

type Logger interface {
	Printf(format string, v ...any)
}
type Encoder interface {
	Encode(any) error
}

type Aggregator struct {
	logger    Logger
	encoder   Encoder
	fs        fs.FS
	aggregate *Stats
	failures  int
}

func NewAggregator(logger Logger, encoder Encoder, fs fs.FS) *Aggregator {
	return &Aggregator{
		logger:    logger,
		encoder:   encoder,
		fs:        fs,
		aggregate: NewStats(""),
	}
}

func (this *Aggregator) Aggregate(args ...string) int {
	for _, arg := range args {
		matches, err := fs.Glob(this.fs, arg)
		if err != nil {
			this.failures++
			continue
		}
		for _, name := range matches {
			info, _ := fs.Stat(this.fs, name)
			if info != nil && info.IsDir() {
				continue
			}
			file, err := this.fs.Open(name)
			if err != nil {
				this.failures++
				continue
			}
			counts := NewStats(name)
			_, err = counts.ReadFrom(file)
			if err != nil {
				this.failures++
			}
			_ = file.Close()

			this.aggregate.Include(counts)
			_ = this.encoder.Encode(counts)
		}
	}
	_ = this.encoder.Encode(this.aggregate)
	return this.failures
}
