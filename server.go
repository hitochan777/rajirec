package rajirec

import (
	"flag"
	"context"
	"github.com/google/subcommands"
)

type ServerCmd struct {
	port int
}

func (*ServerCmd) Name() string {
	return "server"
}

func (*ServerCmd) Synopsis() string {
	return "Run server to record"
}

func (*ServerCmd) Usage() string {
	return "TODO: usage"
}

func (s *ServerCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&s.port, "port", 8080, "port")
}

func (s *ServerCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {

	return subcommands.ExitSuccess
}
