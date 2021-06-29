package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

var data = `
a: 10086
b: help
c:
  d: 1
  e: god
`

type Node struct {
	A int
	B string
	C struct {
		D int    `yaml:"c"`
		E string `yaml:",flow"`
	}
}

func main() {
	node := Node{}
	err := yaml.Unmarshal([]byte(data), &node)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%v\n", node)
	spew.Dump(node)

}
