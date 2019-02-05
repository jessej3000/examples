package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/jessej3000/examples/testgrpc/todo"
)

func main() {
	var err error

	flag.NArg()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}

	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("Unknown subcommand %s", cmd)
	}

	fmt.Println("TODO")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const dbPath = "mydb.pb"

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}

	fmt.Println(proto.MarshalTextString(task))
	return nil
}

func list() error {
	return nil
}
