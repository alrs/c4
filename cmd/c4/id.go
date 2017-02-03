package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/etcenter/c4"
	"golang.org/x/crypto/ssh/terminal"
)

func encode(src io.Reader) *c4.ID {
	id, err := c4.Identify(src)
	if err != nil {
		panic(err)
	}
	return id
}

func fileID(path *string) (*c4.ID, error) {
	f, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	id, err := c4.Identify(f)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func nullId() *c4.ID {
	e := c4.NewIDEncoder()
	io.Copy(e, strings.NewReader(``))
	return e.ID()
}

func printID(id *c4.ID) {
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		fmt.Printf("%s\n", id.String())
	} else {
		fmt.Printf("%s", id.String())
	}
}
