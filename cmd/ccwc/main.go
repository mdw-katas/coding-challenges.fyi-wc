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
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])

	paths := flags.Args()
	encoder := json.NewEncoder(os.Stdout)
	if len(paths) == 0 {
		counts := wc.NewCounts("stdin")
		_, err := counts.ReadFrom(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		err = encoder.Encode(counts)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	aggregate := wc.NewCounts("")
	for _, path := range paths {
		counts := wc.NewCounts(path)
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		_, err = counts.ReadFrom(file)
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = encoder.Encode(counts)
		if err != nil {
			log.Fatal(err)
		}
		aggregate.Include(counts)
	}
	if aggregate.Files > 1 {
		err := encoder.Encode(aggregate)
		if err != nil {
			log.Fatal(err)
		}
	}
}
