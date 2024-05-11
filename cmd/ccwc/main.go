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
	_ = flags.Parse(os.Args[1:])

	paths := flags.Args()
	if len(paths) == 0 {
		counts := wc.NewCounts("stdin")
		_, err := counts.ReadFrom(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		err = json.NewEncoder(os.Stdout).Encode(counts)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	failed := 0
	aggregate := wc.NewCounts("")
	for _, pattern := range paths {
		matches, _ := filepath.Glob(pattern)
		for _, path := range matches {
			info, _ := os.Stat(path)
			if info != nil && info.IsDir() {
				continue
			}
			counts := wc.NewCounts(path)
			ok := countFile(path, counts)
			if !ok {
				failed++
			}
			aggregate.Include(counts)
		}
	}
	if aggregate.Files > 1 {
		err := json.NewEncoder(os.Stdout).Encode(aggregate)
		if err != nil {
			log.Fatal(err)
		}
	}
	if failed > 0 {
		os.Exit(1)
	}
}

func countFile(path string, counts *wc.Counts) bool {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = counts.ReadFrom(file)
	if err != nil {
		log.Println(err)
		return false
	}
	err = file.Close()
	if err != nil {
		log.Println(err)
		return false
	}
	err = json.NewEncoder(os.Stdout).Encode(counts)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
