package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mdwhatcott/coding-challenges.fyi-wc/wc"
)

var Version = "dev"

func main() {
	log.SetFlags(0)

	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage of %s:\n", flags.Name())
		_, _ = fmt.Fprintln(flags.Output(), "ccwc [file ...]")
		_, _ = fmt.Fprintln(flags.Output(), "A tool similar to the built-in `wc` program, but without all the flags because the output is JSON objects.")
		_, _ = fmt.Fprintln(flags.Output(), "If the only argument supplied is '-', read from stdin.")
		_, _ = fmt.Fprintln(flags.Output(), "Note: this tool is quarantined to files rooted below the current working directory.")
		flags.PrintDefaults()
	}
	_ = flags.Parse(os.Args[1:])
	args := flags.Args()
	encoder := json.NewEncoder(os.Stdout)

	if len(args) == 0 {
		log.Println("Error: At least one argument is required.")
		flags.Usage()
		os.Exit(1)
	}
	if len(args) == 1 && args[0] == "-" {
		stats := wc.NewStats("stdin")
		_, _ = stats.ReadFrom(os.Stdin)
		_ = encoder.Encode(stats)
		return
	}
	aggregator := wc.NewAggregator(
		log.Default(),
		encoder,
		os.DirFS("."),
	)
	os.Exit(min(aggregator.Aggregate(args...), 1))
}
