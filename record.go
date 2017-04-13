package rajirec

import (
	"os/exec"
	"net/url"
	"log"
	"path/filepath"
	"strings"
)

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
