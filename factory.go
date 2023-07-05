package clipty

import (
	"github.com/sorenisanerd/gotty/server"
)

type Factory struct {
	persistParams map[string][]string
	mf            MainFunc
}

func NewFactory(persistParams map[string][]string, mf MainFunc) (*Factory, error) {
	return &Factory{
		persistParams: persistParams,
		mf:            mf,
	}, nil
}

func (factory *Factory) Name() string {
	return "cli"
}

func (factory *Factory) New(params map[string][]string) (server.Slave, error) {
	newParams := make(map[string][]string)
	for k, v := range factory.persistParams {
		newParams[k] = v
	}
	for k, v := range params {
		newParams[k] = v
	}
	return New(newParams, factory.mf)
}
