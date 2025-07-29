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
			this.logger.Printf("error globbing arg: %s", arg)
			this.failures++
			continue
		}
		if len(matches) == 0 {
			this.logger.Printf("no results found for arg: %s", arg)
			this.failures++
			continue
		}
		for _, name := range matches {
			info, err := fs.Stat(this.fs, name)
			if info != nil && info.IsDir() {
				continue
			}
			if err != nil {
				this.logger.Printf("WARNING: failed to stat %s: %s", name, err)
				this.failures++
				continue
			}
			file, err := this.fs.Open(name)
			if err != nil {
				this.logger.Printf("error opening file %s: %s", name, err)
				this.failures++
				continue
			}
			counts := NewStats(name)
			_, err = counts.ReadFrom(file)
			if err != nil {
				this.logger.Printf("error reading file %s: %s", name, err)
				this.failures++
			}
			_ = file.Close()

			this.aggregate.Include(counts)
			_ = this.encoder.Encode(counts)
		}
	}
	if this.aggregate.Files > 0 {
		_ = this.encoder.Encode(this.aggregate)
	}
	return this.failures
}
