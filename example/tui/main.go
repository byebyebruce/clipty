package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/byebyebruce/clipty"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
	"github.com/sorenisanerd/gotty/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opt := &server.Options{
		PermitArguments: true,
		PermitWrite:     true,
		TitleVariables: map[string]interface{}{
			"command": "my cli",
		},
	}
	persistParams := map[string][]string{"mode": {"debug"}}
	err := clipty.RunServer(ctx, opt, persistParams, mainFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func mainFunc(ctx context.Context, params map[string][]string, stdin *os.File, stdout *os.File, stderr *os.File) {
	fmt.Fprintln(stdout, color.GreenString("Welcome to my cli"))
	fmt.Fprintln(stdout, "params", params)

	for {
		// prompt
		pt := promptui.Prompt{
			Stdin:   stdin,
			Stdout:  stdout,
			Label:   "Input something",
			Default: "xxxx",
		}
		str, err := pt.Run()
		if err != nil {
			fmt.Fprint(stderr, color.GreenString("error "+err.Error()))
			return
		}
		str = strings.TrimSpace(str)
		_, err = fmt.Fprintln(stdout, color.GreenString("prompt "+str))

		// table
		table := tablewriter.NewWriter(stdout)
		table.SetHeader([]string{"#", "CMD", "DESC"})
		table.SetAutoWrapText(false)
		for i, c := range []string{"a", "b", "c"} {
			table.Append([]string{strconv.Itoa(i + 1), color.GreenString(c), "desc" + c})
		}
		table.Render()

		// progressbar
		bar := progressbar.NewOptions(100,
			progressbar.OptionSetWriter(stdout),
			progressbar.OptionShowCount(),
			//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionEnableColorCodes(true),
			//progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(15),
			progressbar.OptionSetDescription("[cyan]"+"hello"+"..."),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[green]=[reset]",
				SaucerHead:    "[green]>[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}))
		for i := 0; i < 100; i++ {
			bar.Add(1)
			time.Sleep(10 * time.Millisecond)
		}
		bar.Exit()
	}
}
