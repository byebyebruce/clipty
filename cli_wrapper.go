package clipty

import (
	"context"
	"os"

	"github.com/creack/pty"
)

// MainFunc is CLI main loop  function.
type MainFunc func(ctx context.Context, params map[string][]string, stdin *os.File, stdout *os.File, stderr *os.File)

type CliWrapper struct {
	params map[string][]string
	pty    *os.File
	tty    *os.File
	cancel context.CancelFunc
}

func New(params map[string][]string, mf MainFunc) (*CliWrapper, error) {
	pty, tty, err := pty.Open()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := &CliWrapper{
		params: params,
		pty:    pty,
		tty:    tty,
		cancel: cancel,
	}

	// When the process is closed by the user,
	// close pty so that Read() on the pty breaks with an EOF.
	go func() {
		defer func() {
			tty.Sync()
			c.Close()
		}()

		mf(ctx, params, tty, tty, tty)
	}()

	return c, nil
}

func (c *CliWrapper) Read(p []byte) (n int, err error) {
	return c.pty.Read(p)
}

func (c *CliWrapper) Write(p []byte) (n int, err error) {
	return c.pty.Write(p)
}

func (c *CliWrapper) Close() error {
	c.cancel()
	c.tty.Close()
	return c.pty.Close()
}

func (c *CliWrapper) WindowTitleVariables() map[string]interface{} {
	return map[string]interface{}{
		//"command": c.command,
		"params": c.params,
	}
}

func (c *CliWrapper) ResizeTerminal(width int, height int) error {
	window := pty.Winsize{
		Rows: uint16(height),
		Cols: uint16(width),
		X:    0,
		Y:    0,
	}
	err := pty.Setsize(c.pty, &window)
	if err != nil {
		return err
	} else {
		return nil
	}
}
