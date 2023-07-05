package clipty

import (
	"context"

	"github.com/sorenisanerd/gotty/server"
)

func RunServer(ctx context.Context, opt *server.Options, persistParams map[string][]string, mainFunc MainFunc) error {
	if len(opt.Address) == 0 {
		opt.Address = "0.0.0.0"
	}
	if len(opt.Port) == 0 {
		opt.Port = "8081"
	}

	if err := opt.Validate(); err != nil {
		return err
	}

	factory, err := NewFactory(persistParams, mainFunc)
	if err != nil {
		return err
	}

	srv, err := server.New(factory, opt)
	if err != nil {
		return err
	}

	err = srv.Run(ctx)
	if err != nil && err == context.Canceled {
		return nil
	}

	return err
}
