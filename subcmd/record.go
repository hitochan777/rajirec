package subcmd

import (
	"os/exec"
	"net/url"
	"log"
	"path/filepath"
	"strings"
	"flag"
	"context"
	"strconv"

	"github.com/google/subcommands"
)

type RecordCmd struct {
	duration int
	stationId string
	outputFile string
	channel string
}

func (*RecordCmd) Name() string { return "record" }
func (*RecordCmd) Synopsis() string { return "record live stream" }
func (*RecordCmd) Usage() string {
	return "Record live stream\n"
}
func (r *RecordCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&r.duration, "duration", 0, "duration of recording(min)")
	f.StringVar(&r.stationId, "areaid", "", "Station ID")
	f.StringVar(&r.outputFile, "output", "", "path to output file")
	f.StringVar(&r.channel, "channel", "fm", "Channel")
}

func (r *RecordCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	config := NewConfig()
	log.Println("recording...")
	areas := NewAreas(config.General.API_URL)
	if streamURL, err := areas.GetStreamURL(r.stationId, r.channel); err != nil {
		log.Fatal(err)
	} else {
		Record(streamURL, r.outputFile, r.duration)
	}
	return subcommands.ExitSuccess
}

func Record(streamURL string, outputPath string, duration int) {
	rtmpdumpPath, err := exec.LookPath("rtmpdump")
	if err != nil {
		log.Fatal("rtmpdump is not installed. Please install it.")
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
		"--stop", strconv.Itoa(duration * 60),
		"--live",
		"-o", outputPath,
	).Run()
}
