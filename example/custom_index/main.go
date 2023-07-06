package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/byebyebruce/clipty"
	"github.com/fatih/color"
	"github.com/sorenisanerd/gotty/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opt := &server.Options{
		PermitWrite: true,
		TitleVariables: map[string]interface{}{
			"command": "stdin",
		},
		IndexFile: "index.html",
	}
	err := clipty.RunServer(ctx, opt, nil, mainFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func mainFunc(ctx context.Context, params map[string][]string, stdin *os.File, stdout *os.File, stderr *os.File) {
	r := bufio.NewReader(stdin)
	for {
		fmt.Fprintln(stdout, color.GreenString("Please input your name"))
		// some input
		n, err := r.ReadString('\n')
		if err != nil {
			return
		}
		n = strings.TrimSpace(n)
		_, err = fmt.Fprintln(stdout, color.GreenString("Hello, "+n))
		if err != nil {
			return
		}
	}
}
