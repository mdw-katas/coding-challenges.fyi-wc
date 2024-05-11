# The Coding Challenges - Build Your Own `wc` Tool

https://codingchallenges.fyi/challenges/challenge-wc/

Example output (of this project's root directory):

```
github.com/mdwhatcott/coding-challenges.fyi-wc
$ make install && ccwc *
go version go1.22.0 darwin/amd64
go test -race -cover -timeout=1s -count=1 ./...
?       github.com/mdwhatcott/coding-challenges.fyi-wc	[no test files]
	    github.com/mdwhatcott/coding-challenges.fyi-wc/cmd/ccwc		coverage: 0.0% of statements
ok      github.com/mdwhatcott/coding-challenges.fyi-wc/wc	0.268s	coverage: 85.3% of statements
go install -ldflags="-X 'main.Version=v1.0.0'" github.com/mdwhatcott/coding-challenges.fyi-wc/cmd/...
{"name":"LICENSE.md","lines":21,"words":170,"bytes":1075}
{"name":"Makefile","lines":12,"words":34,"bytes":285}
{"name":"README.md","lines":61,"words":131,"bytes":1221}
{"name":"deps.go","lines":9,"words":19,"bytes":204}
{"name":"go.mod","lines":9,"words":13,"bytes":190}
{"name":"go.sum","lines":6,"words":18,"bytes":531}
{"files":6,"lines":118,"words":385,"bytes":3506}
```

Pipe output to `jq -s` to wrap in JSON array:

```
github.com/mdwhatcott/coding-challenges.fyi-wc
$ make install && ccwc * | jq -s .
go version go1.22.0 darwin/amd64
go test -race -cover -timeout=1s -count=1 ./...
?       github.com/mdwhatcott/coding-challenges.fyi-wc	[no test files]
	    github.com/mdwhatcott/coding-challenges.fyi-wc/cmd/ccwc		coverage: 0.0% of statements
ok      github.com/mdwhatcott/coding-challenges.fyi-wc/wc	0.271s	coverage: 85.3% of statements
go install -ldflags="-X 'main.Version=v1.0.0'" github.com/mdwhatcott/coding-challenges.fyi-wc/cmd/...
[
  {
    "name": "LICENSE.md",
    "lines": 21,
    "words": 170,
    "bytes": 1075
  },
  {
    "name": "Makefile",
    "lines": 12,
    "words": 34,
    "bytes": 285
  },
  {
    "name": "README.md",
    "lines": 58,
    "words": 133,
    "bytes": 1229
  },
  {
    "name": "deps.go",
    "lines": 9,
    "words": 19,
    "bytes": 204
  },
  {
    "name": "go.mod",
    "lines": 9,
    "words": 13,
    "bytes": 190
  },
  {
    "name": "go.sum",
    "lines": 6,
    "words": 18,
    "bytes": 531
  },
  {
    "files": 6,
    "lines": 115,
    "words": 387,
    "bytes": 3514
  }
]
```
