package command

import (
	"flag"
	"net/http"

	"github.com/mitchellh/cli"
	clihttp "github.com/patdhlk/cli/http"
)

var (
	addr = flag.String("http", ":8080", "http listen address")
)

type ServerCommand struct {
	ShutdownCh <-chan struct{}
	Ui         cli.Ui
	Meta
}

func (c *ServerCommand) Run(args []string) int {
	flag.Parse()
	//Initialize your HTTP Server
	server := &http.Server{}
	server.Handler = clihttp.Handler()
	server.Addr = *addr
	c.Ui.Output("==> starting server")
	go server.ListenAndServe()

	c.Ui.Output("")
	c.Ui.Output("==> MYAPP Server started! Log data will stream below")

	//wait for shutdown
	select {
	case <-c.ShutdownCh:
		c.Ui.Output("==> MYAPP shutdown triggered")
	}
	return 0
}

func (c *ServerCommand) Synopsis() string {
	return "Start a MYAPP Server"
}

func (c *ServerCommand) Help() string {
	return ""
}
