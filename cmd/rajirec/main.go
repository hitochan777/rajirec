package main

import (
	"log"

	"github.com/hitochan777/rajirec"
)

func main(){
	log.Println("recording...")
	rajirec.Record("rtmpe://netradio-r2-flash.nhk.jp/live/NetRadio_R2_flash@63342", "hoge.m4a")
}
