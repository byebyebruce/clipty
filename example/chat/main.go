package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/byebyebruce/clipty"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
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
	persistParams := map[string][]string{"who": {"bruce"}}
	err := clipty.RunServer(ctx, opt, persistParams, mainFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func mainFunc(ctx context.Context, params map[string][]string, stdin *os.File, stdout *os.File, stderr *os.File) {
	fmt.Fprintln(stdout, color.GreenString("Hello, "+params["who"][0]))

	p := tea.NewProgram(initialModel(), tea.WithContext(ctx), tea.WithAltScreen(), tea.WithInput(stdin), tea.WithOutput(stdout))

	if _, err := p.Run(); err != nil {
		fmt.Fprintln(stdout, color.RedString("error: %v", err))
	}
}
