package rajirec

import (
	"os/exec"
	"net/url"
	"log"
	"path/filepath"
	"strings"
	"flag"
	"context"
	"github.com/google/subcommands"
)

type RecordCmd struct {
	duration string
	stationId string
	outputFile string
}

func (*RecordCmd) Name() string { return "record" }
func (*RecordCmd) Synopsis() string { return "record live stream" }
func (*RecordCmd) Usage() string {return "hoge"}
func (r *RecordCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&r.duration, "duration", "", "duration of recording")
	f.StringVar(&r.stationId, "sid", "", "Station ID")
	f.StringVar(&r.outputFile, "output", "", "path to output file")
}

func (r *RecordCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	log.Println("recording...")
	Record("rtmpe://netradio-r2-flash.nhk.jp/live/NetRadio_R2_flash@63342", "hoge.m4a")
	return subcommands.ExitSuccess
}

func Record(streamURL string, outputPath string) {
	rtmpdumpPath, err := exec.LookPath("rtmpdump")
	if err != nil {
		log.Fatal("rtmpdump is not installed")
	}
	u, err := url.Parse(streamURL)
	rtmp := u.Scheme + "://" + u.Host
	app, playPath := filepath.Split(u.RequestURI())
	app = strings.Trim(app, "/")
	exec.Command(
		rtmpdumpPath,
		"--rtmp", rtmp,
		"--playpath", playPath,
		"--swfVfy", "http://www3.nhk.or.jp/netradio/files/swf/rtmpe.swf",
		"--app", app,
		"--live",
		"-o", outputPath,
	).Run()
}
