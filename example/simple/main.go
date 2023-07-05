package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/byebyebruce/clipty"
	"github.com/sorenisanerd/gotty/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	opt := &server.Options{}
	err := clipty.RunServer(ctx, opt, nil, mainFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func mainFunc(ctx context.Context, params map[string][]string, stdin *os.File, stdout *os.File, stderr *os.File) {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
		fmt.Fprintln(stdout, "sleep", i+1)
	}
	fmt.Fprintln(stdout, "Bye..")
}
